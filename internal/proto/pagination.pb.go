// Copyright 2022 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/pagination.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Cursor represents the position of a paginated Bigtable cache data for an
// import group.
type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The index of the import group, starts from 0.
	ImportGroup int32 `protobuf:"varint,1,opt,name=import_group,json=importGroup,proto3" json:"import_group,omitempty"`
	// The index of the page, starts from 0.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// The position of the next read item in the current page, starts from 0.
	Item int32 `protobuf:"varint,3,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *Cursor) GetImportGroup() int32 {
	if x != nil {
		return x.ImportGroup
	}
	return 0
}

func (x *Cursor) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Cursor) GetItem() int32 {
	if x != nil {
		return x.Item
	}
	return 0
}

// Represents cursors of several import groups. This holds the position
// information of one entity.
type CursorGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entity DCID or other information that identifies the CursorGroup.
	Key     string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Cursors []*Cursor `protobuf:"bytes,2,rep,name=cursors,proto3" json:"cursors,omitempty"`
}

func (x *CursorGroup) Reset() {
	*x = CursorGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CursorGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CursorGroup) ProtoMessage() {}

func (x *CursorGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CursorGroup.ProtoReflect.Descriptor instead.
func (*CursorGroup) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *CursorGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CursorGroup) GetCursors() []*Cursor {
	if x != nil {
		return x.Cursors
	}
	return nil
}

// Represents the cursor information of one pagination request.
type PaginationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each cursor group corresponds to the cursor information of one requested
	// entity. There are multiple cursor groups for bulk APIs.
	CursorGroups []*CursorGroup `protobuf:"bytes,1,rep,name=cursor_groups,json=cursorGroups,proto3" json:"cursor_groups,omitempty"`
}

func (x *PaginationInfo) Reset() {
	*x = PaginationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationInfo) ProtoMessage() {}

func (x *PaginationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationInfo.ProtoReflect.Descriptor instead.
func (*PaginationInfo) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationInfo) GetCursorGroups() []*CursorGroup {
	if x != nil {
		return x.CursorGroups
	}
	return nil
}

var File_v1_pagination_proto protoreflect.FileDescriptor

var file_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x53, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x51, 0x0a, 0x0b, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x73, 0x22, 0x52, 0x0a,
	0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x40, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_pagination_proto_rawDescOnce sync.Once
	file_v1_pagination_proto_rawDescData = file_v1_pagination_proto_rawDesc
)

func file_v1_pagination_proto_rawDescGZIP() []byte {
	file_v1_pagination_proto_rawDescOnce.Do(func() {
		file_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pagination_proto_rawDescData)
	})
	return file_v1_pagination_proto_rawDescData
}

var file_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_pagination_proto_goTypes = []interface{}{
	(*Cursor)(nil),         // 0: datacommons.v1.Cursor
	(*CursorGroup)(nil),    // 1: datacommons.v1.CursorGroup
	(*PaginationInfo)(nil), // 2: datacommons.v1.PaginationInfo
}
var file_v1_pagination_proto_depIdxs = []int32{
	0, // 0: datacommons.v1.CursorGroup.cursors:type_name -> datacommons.v1.Cursor
	1, // 1: datacommons.v1.PaginationInfo.cursor_groups:type_name -> datacommons.v1.CursorGroup
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_pagination_proto_init() }
func file_v1_pagination_proto_init() {
	if File_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CursorGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_pagination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_pagination_proto_goTypes,
		DependencyIndexes: file_v1_pagination_proto_depIdxs,
		MessageInfos:      file_v1_pagination_proto_msgTypes,
	}.Build()
	File_v1_pagination_proto = out.File
	file_v1_pagination_proto_rawDesc = nil
	file_v1_pagination_proto_goTypes = nil
	file_v1_pagination_proto_depIdxs = nil
}
