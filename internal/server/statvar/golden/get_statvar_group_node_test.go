// Copyright 2021 Google LLC
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

package golden

import (
	"context"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetStatVarGroupNode(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "get_statvar_group_node")

	testSuite := func(mixer pb.MixerClient, recon pb.ReconClient, latencyTest bool) {
		for _, c := range []struct {
			places     []string
			svg        string
			goldenFile string
		}{
			{
				[]string{"country/USA"},
				"dc/g/Person_EducationalAttainment",
				"school.json",
			},
			{
				[]string{"country/USA"},
				"dc/g/Person_EnrollmentLevel-EnrolledInCollegeUndergraduateYears_Race",
				"school_race.json",
			},
			{
				[]string{"country/GBR"},
				"dc/g/Demographics",
				"demographics_gbr.json",
			},
			// Run this first to test the server cache is not modified.
			{
				[]string{"geoId/0649670"},
				"dc/g/Root",
				"root_mtv.json",
			},
			// Run this first to test the server cache is not modified.
			{
				[]string{"geoId/0649670", "country/JPN"},
				"dc/g/Root",
				"root_mtv_jpn.json",
			},
			{
				[]string{},
				"dc/g/Root",
				"root.json",
			},
			{
				[]string{},
				"dc/g/Demographics",
				"demographics.json",
			},
			{
				[]string{"geoId/0649670"},
				"dc/g/Person_CitizenshipStatus-NotAUSCitizen_CorrectionalFacilityOperator-StateOperated&FederallyOperated&PrivatelyOperated",
				"citizenship.json",
			},
			{
				[]string{"geoId/06"},
				"g/Feeding_America",
				"private.json",
			},
		} {
			resp, err := mixer.GetStatVarGroupNode(ctx, &pb.GetStatVarGroupNodeRequest{
				Places:       c.places,
				StatVarGroup: c.svg,
			})
			if err != nil {
				t.Errorf("could not GetStatVarGroupNode: %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			if test.GenerateGolden {
				test.UpdateProtoGolden(resp, goldenPath, c.goldenFile)
				continue
			}

			var expected pb.StatVarGroupNode
			if err = test.ReadJSON(goldenPath, c.goldenFile, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file")
				continue
			}

			if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("payload got diff: %v", diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"GetStatVarGroupNode",
		&test.TestOption{UseCache: true, UseMemdb: true},
		testSuite,
	); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}

}
