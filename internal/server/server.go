// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/bigtable"
	pubsub "cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/datacommonsorg/mixer/internal/parser/mcf"
	pb "github.com/datacommonsorg/mixer/internal/proto"
	dcpubsub "github.com/datacommonsorg/mixer/internal/pubsub"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/translator/solver"
	"github.com/datacommonsorg/mixer/internal/translator/types"
)

// Cache holds cached data for the mixer server.
type Cache struct {
	// ParentSvg is a map of sv/svg id to a list of its parent svgs sorted alphabetically.
	ParentSvg map[string][]string
	// SvgInfo is a map of svg id to its information.
	SvgInfo                   map[string]*pb.StatVarGroupNode
	SvgSearchIndex            *SearchIndex
	BlocklistedSvgSearchIndex *SearchIndex
}

// Metadata represents the metadata used by the server.
type Metadata struct {
	Mappings         []*types.Mapping
	OutArcInfo       map[string]map[string][]types.OutArcInfo
	InArcInfo        map[string][]types.InArcInfo
	SubTypeMap       map[string]string
	Bq               string
	BtProject        string
	BranchBtInstance string
}

// Server holds resources for a mixer server
type Server struct {
	store    *store.Store
	metadata *Metadata
	cache    *Cache
}

func (s *Server) updateBranchTable(ctx context.Context, branchTableName string) {
	branchTable, err := NewBtTable(
		ctx, s.metadata.BtProject, s.metadata.BranchBtInstance, branchTableName)
	if err != nil {
		log.Printf("Failed to udpate branch cache Bigtable client: %v", err)
		return
	}
	s.store.UpdateBranchBt(branchTable)
}

// ReadBranchTableName reads branch cache folder from GCS.
func ReadBranchTableName(
	ctx context.Context, bucket, versionFile string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	rc, err := client.Bucket(bucket).Object(versionFile).NewReader(ctx)
	if err != nil {
		return "", err
	}
	defer rc.Close()
	folder, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", err
	}
	return string(folder), nil
}

// NewMetadata initialize the metadata for translator.
func NewMetadata(
	bqDataset, storeProject, branchInstance, schemaPath string) (*Metadata, error) {
	_, filename, _, _ := runtime.Caller(0)
	subTypeMap, err := solver.GetSubTypeMap(
		path.Join(path.Dir(filename), "../translator/table_types.json"))
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(schemaPath)
	if err != nil {
		return nil, err
	}
	mappings := []*types.Mapping{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".mcf") {
			mappingStr, err := ioutil.ReadFile(filepath.Join(schemaPath, f.Name()))
			if err != nil {
				return nil, err
			}
			mapping, err := mcf.ParseMapping(string(mappingStr), bqDataset)
			if err != nil {
				return nil, err
			}
			mappings = append(mappings, mapping...)
		}
	}
	outArcInfo := map[string]map[string][]types.OutArcInfo{}
	inArcInfo := map[string][]types.InArcInfo{}
	return &Metadata{
			mappings,
			outArcInfo,
			inArcInfo,
			subTypeMap,
			bqDataset,
			storeProject,
			branchInstance,
		},
		nil
}

// NewBtTable creates a new bigtable.Table instance.
func NewBtTable(
	ctx context.Context, projectID, instanceID, tableID string) (
	*bigtable.Table, error) {
	btClient, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		return nil, err
	}
	return btClient.Open(tableID), nil
}

// SubscribeBranchCacheUpdate subscribe for branch cache update.
func (s *Server) SubscribeBranchCacheUpdate(
	ctx context.Context, pubsubProject, subscriberPrefix, pubsubTopic string,
) error {
	return dcpubsub.Subscribe(
		ctx,
		pubsubProject,
		subscriberPrefix,
		pubsubTopic,
		func(ctx context.Context, msg *pubsub.Message) error {
			branchTableName := string(msg.Data)
			log.Printf("Branch Cache Subscriber: use branch cache %s\n", branchTableName)
			s.updateBranchTable(ctx, branchTableName)
			return nil
		},
	)
}

// NewCache initializes the cache for stat var hierarchy.
func NewCache(ctx context.Context, baseTable *bigtable.Table) (*Cache, error) {
	rawSvg, err := GetRawSvg(ctx, baseTable)
	if err != nil {
		return nil, err
	}
	parentSvgMap := BuildParentSvgMap(rawSvg)
	searchIndex := BuildStatVarSearchIndex(rawSvg, false)
	blocklistedSearchIndex := BuildStatVarSearchIndex(rawSvg, true)

	return &Cache{
		ParentSvg:                 parentSvgMap,
		SvgInfo:                   rawSvg,
		SvgSearchIndex:            searchIndex,
		BlocklistedSvgSearchIndex: blocklistedSearchIndex,
	}, nil
}

// NewServer creates a new server instance.
func NewServer(
	bqClient *bigquery.Client,
	baseTable *bigtable.Table,
	branchTable *bigtable.Table,
	metadata *Metadata,
	cache *Cache,
	memdb *store.MemDb,
) *Server {
	return &Server{
		store:    store.NewStore(bqClient, memdb, baseTable, branchTable),
		metadata: metadata,
		cache:    cache,
	}
}
