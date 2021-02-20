//
// Copyright (C) 2021 Grakn Labs
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: v2/protobuf/cluster/cluster.proto

package grakn_protocol_cluster

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_proto_rawDescGZIP(), []int{0}
}

type Cluster_Servers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cluster_Servers) Reset() {
	*x = Cluster_Servers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster_Servers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster_Servers) ProtoMessage() {}

func (x *Cluster_Servers) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster_Servers.ProtoReflect.Descriptor instead.
func (*Cluster_Servers) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_proto_rawDescGZIP(), []int{0, 0}
}

type Cluster_Servers_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cluster_Servers_Req) Reset() {
	*x = Cluster_Servers_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster_Servers_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster_Servers_Req) ProtoMessage() {}

func (x *Cluster_Servers_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster_Servers_Req.ProtoReflect.Descriptor instead.
func (*Cluster_Servers_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Cluster_Servers_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []string `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *Cluster_Servers_Res) Reset() {
	*x = Cluster_Servers_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster_Servers_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster_Servers_Res) ProtoMessage() {}

func (x *Cluster_Servers_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster_Servers_Res.ProtoReflect.Descriptor instead.
func (*Cluster_Servers_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Cluster_Servers_Res) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

var File_v2_protobuf_cluster_cluster_proto protoreflect.FileDescriptor

var file_v2_protobuf_cluster_cluster_proto_rawDesc = []byte{
	0x0a, 0x21, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x07, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x31, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x42, 0x5a, 0x0a, 0x16, 0x67, 0x72, 0x61,
	0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x32, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3b, 0x67,
	0x72, 0x61, 0x6b, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_cluster_cluster_proto_rawDescOnce sync.Once
	file_v2_protobuf_cluster_cluster_proto_rawDescData = file_v2_protobuf_cluster_cluster_proto_rawDesc
)

func file_v2_protobuf_cluster_cluster_proto_rawDescGZIP() []byte {
	file_v2_protobuf_cluster_cluster_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_cluster_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_cluster_cluster_proto_rawDescData)
	})
	return file_v2_protobuf_cluster_cluster_proto_rawDescData
}

var file_v2_protobuf_cluster_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v2_protobuf_cluster_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil),             // 0: grakn.protocol.cluster.Cluster
	(*Cluster_Servers)(nil),     // 1: grakn.protocol.cluster.Cluster.Servers
	(*Cluster_Servers_Req)(nil), // 2: grakn.protocol.cluster.Cluster.Servers.Req
	(*Cluster_Servers_Res)(nil), // 3: grakn.protocol.cluster.Cluster.Servers.Res
}
var file_v2_protobuf_cluster_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2_protobuf_cluster_cluster_proto_init() }
func file_v2_protobuf_cluster_cluster_proto_init() {
	if File_v2_protobuf_cluster_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_cluster_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_v2_protobuf_cluster_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster_Servers); i {
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
		file_v2_protobuf_cluster_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster_Servers_Req); i {
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
		file_v2_protobuf_cluster_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster_Servers_Res); i {
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
			RawDescriptor: file_v2_protobuf_cluster_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_cluster_cluster_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_cluster_cluster_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_cluster_cluster_proto_msgTypes,
	}.Build()
	File_v2_protobuf_cluster_cluster_proto = out.File
	file_v2_protobuf_cluster_cluster_proto_rawDesc = nil
	file_v2_protobuf_cluster_cluster_proto_goTypes = nil
	file_v2_protobuf_cluster_cluster_proto_depIdxs = nil
}
