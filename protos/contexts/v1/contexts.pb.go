//
//  Copyright 2020 Docker Compose CLI authors

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: protos/contexts/v1/contexts.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContextType string `protobuf:"bytes,2,opt,name=contextType,proto3" json:"contextType,omitempty"`
	Current     bool   `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_contexts_v1_contexts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_protos_contexts_v1_contexts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_protos_contexts_v1_contexts_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Context) GetContextType() string {
	if x != nil {
		return x.ContextType
	}
	return ""
}

func (x *Context) GetCurrent() bool {
	if x != nil {
		return x.Current
	}
	return false
}

type SetCurrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetCurrentRequest) Reset() {
	*x = SetCurrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_contexts_v1_contexts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentRequest) ProtoMessage() {}

func (x *SetCurrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_contexts_v1_contexts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentRequest.ProtoReflect.Descriptor instead.
func (*SetCurrentRequest) Descriptor() ([]byte, []int) {
	return file_protos_contexts_v1_contexts_proto_rawDescGZIP(), []int{1}
}

func (x *SetCurrentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetCurrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetCurrentResponse) Reset() {
	*x = SetCurrentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_contexts_v1_contexts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCurrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCurrentResponse) ProtoMessage() {}

func (x *SetCurrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_contexts_v1_contexts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCurrentResponse.ProtoReflect.Descriptor instead.
func (*SetCurrentResponse) Descriptor() ([]byte, []int) {
	return file_protos_contexts_v1_contexts_proto_rawDescGZIP(), []int{2}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_contexts_v1_contexts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_contexts_v1_contexts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_protos_contexts_v1_contexts_proto_rawDescGZIP(), []int{3}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contexts []*Context `protobuf:"bytes,1,rep,name=contexts,proto3" json:"contexts,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_contexts_v1_contexts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_contexts_v1_contexts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_protos_contexts_v1_contexts_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetContexts() []*Context {
	if x != nil {
		return x.Contexts
	}
	return nil
}

var File_protos_contexts_v1_contexts_proto protoreflect.FileDescriptor

var file_protos_contexts_v1_contexts_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x59, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x27, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x55,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x73, 0x32, 0xea, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2d,
	0x63, 0x6c, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_contexts_v1_contexts_proto_rawDescOnce sync.Once
	file_protos_contexts_v1_contexts_proto_rawDescData = file_protos_contexts_v1_contexts_proto_rawDesc
)

func file_protos_contexts_v1_contexts_proto_rawDescGZIP() []byte {
	file_protos_contexts_v1_contexts_proto_rawDescOnce.Do(func() {
		file_protos_contexts_v1_contexts_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_contexts_v1_contexts_proto_rawDescData)
	})
	return file_protos_contexts_v1_contexts_proto_rawDescData
}

var file_protos_contexts_v1_contexts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_contexts_v1_contexts_proto_goTypes = []interface{}{
	(*Context)(nil),            // 0: com.docker.api.protos.context.v1.Context
	(*SetCurrentRequest)(nil),  // 1: com.docker.api.protos.context.v1.SetCurrentRequest
	(*SetCurrentResponse)(nil), // 2: com.docker.api.protos.context.v1.SetCurrentResponse
	(*ListRequest)(nil),        // 3: com.docker.api.protos.context.v1.ListRequest
	(*ListResponse)(nil),       // 4: com.docker.api.protos.context.v1.ListResponse
}
var file_protos_contexts_v1_contexts_proto_depIdxs = []int32{
	0, // 0: com.docker.api.protos.context.v1.ListResponse.contexts:type_name -> com.docker.api.protos.context.v1.Context
	1, // 1: com.docker.api.protos.context.v1.Contexts.SetCurrent:input_type -> com.docker.api.protos.context.v1.SetCurrentRequest
	3, // 2: com.docker.api.protos.context.v1.Contexts.List:input_type -> com.docker.api.protos.context.v1.ListRequest
	2, // 3: com.docker.api.protos.context.v1.Contexts.SetCurrent:output_type -> com.docker.api.protos.context.v1.SetCurrentResponse
	4, // 4: com.docker.api.protos.context.v1.Contexts.List:output_type -> com.docker.api.protos.context.v1.ListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_contexts_v1_contexts_proto_init() }
func file_protos_contexts_v1_contexts_proto_init() {
	if File_protos_contexts_v1_contexts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_contexts_v1_contexts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_protos_contexts_v1_contexts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentRequest); i {
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
		file_protos_contexts_v1_contexts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCurrentResponse); i {
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
		file_protos_contexts_v1_contexts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_protos_contexts_v1_contexts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_protos_contexts_v1_contexts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_contexts_v1_contexts_proto_goTypes,
		DependencyIndexes: file_protos_contexts_v1_contexts_proto_depIdxs,
		MessageInfos:      file_protos_contexts_v1_contexts_proto_msgTypes,
	}.Build()
	File_protos_contexts_v1_contexts_proto = out.File
	file_protos_contexts_v1_contexts_proto_rawDesc = nil
	file_protos_contexts_v1_contexts_proto_goTypes = nil
	file_protos_contexts_v1_contexts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContextsClient is the client API for Contexts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContextsClient interface {
	// Sets the current request for all calls
	SetCurrent(ctx context.Context, in *SetCurrentRequest, opts ...grpc.CallOption) (*SetCurrentResponse, error)
	// Returns the list of existing contexts
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type contextsClient struct {
	cc grpc.ClientConnInterface
}

func NewContextsClient(cc grpc.ClientConnInterface) ContextsClient {
	return &contextsClient{cc}
}

func (c *contextsClient) SetCurrent(ctx context.Context, in *SetCurrentRequest, opts ...grpc.CallOption) (*SetCurrentResponse, error) {
	out := new(SetCurrentResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.context.v1.Contexts/SetCurrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/com.docker.api.protos.context.v1.Contexts/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextsServer is the server API for Contexts service.
type ContextsServer interface {
	// Sets the current request for all calls
	SetCurrent(context.Context, *SetCurrentRequest) (*SetCurrentResponse, error)
	// Returns the list of existing contexts
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedContextsServer can be embedded to have forward compatible implementations.
type UnimplementedContextsServer struct {
}

func (*UnimplementedContextsServer) SetCurrent(context.Context, *SetCurrentRequest) (*SetCurrentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCurrent not implemented")
}
func (*UnimplementedContextsServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterContextsServer(s *grpc.Server, srv ContextsServer) {
	s.RegisterService(&_Contexts_serviceDesc, srv)
}

func _Contexts_SetCurrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCurrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).SetCurrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.context.v1.Contexts/SetCurrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).SetCurrent(ctx, req.(*SetCurrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contexts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.docker.api.protos.context.v1.Contexts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contexts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.docker.api.protos.context.v1.Contexts",
	HandlerType: (*ContextsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetCurrent",
			Handler:    _Contexts_SetCurrent_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Contexts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/contexts/v1/contexts.proto",
}
