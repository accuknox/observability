// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: src/proto/summary/summary.proto

package summary

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

type LogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label     string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *LogsRequest) Reset() {
	*x = LogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_summary_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRequest) ProtoMessage() {}

func (x *LogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_summary_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRequest.ProtoReflect.Descriptor instead.
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_summary_summary_proto_rawDescGZIP(), []int{0}
}

func (x *LogsRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LogsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type LogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodDetail     string            `protobuf:"bytes,1,opt,name=podDetail,proto3" json:"podDetail,omitempty"`
	Namespace     string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ListOfProcess []*ListOfSource   `protobuf:"bytes,3,rep,name=listOfProcess,proto3" json:"listOfProcess,omitempty"`
	ListOfFile    []*ListOfSource   `protobuf:"bytes,4,rep,name=listOfFile,proto3" json:"listOfFile,omitempty"`
	ListOfNetwork []*ListOfSource   `protobuf:"bytes,5,rep,name=listOfNetwork,proto3" json:"listOfNetwork,omitempty"`
	Ingress       *ListOfConnection `protobuf:"bytes,6,opt,name=ingress,proto3" json:"ingress,omitempty"`
	Egress        *ListOfConnection `protobuf:"bytes,7,opt,name=egress,proto3" json:"egress,omitempty"`
}

func (x *LogsResponse) Reset() {
	*x = LogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_summary_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponse) ProtoMessage() {}

func (x *LogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_summary_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponse.ProtoReflect.Descriptor instead.
func (*LogsResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_summary_summary_proto_rawDescGZIP(), []int{1}
}

func (x *LogsResponse) GetPodDetail() string {
	if x != nil {
		return x.PodDetail
	}
	return ""
}

func (x *LogsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *LogsResponse) GetListOfProcess() []*ListOfSource {
	if x != nil {
		return x.ListOfProcess
	}
	return nil
}

func (x *LogsResponse) GetListOfFile() []*ListOfSource {
	if x != nil {
		return x.ListOfFile
	}
	return nil
}

func (x *LogsResponse) GetListOfNetwork() []*ListOfSource {
	if x != nil {
		return x.ListOfNetwork
	}
	return nil
}

func (x *LogsResponse) GetIngress() *ListOfConnection {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *LogsResponse) GetEgress() *ListOfConnection {
	if x != nil {
		return x.Egress
	}
	return nil
}

type ListOfSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source   string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Resource []string `protobuf:"bytes,2,rep,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ListOfSource) Reset() {
	*x = ListOfSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_summary_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfSource) ProtoMessage() {}

func (x *ListOfSource) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_summary_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfSource.ProtoReflect.Descriptor instead.
func (*ListOfSource) Descriptor() ([]byte, []int) {
	return file_src_proto_summary_summary_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfSource) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ListOfSource) GetResource() []string {
	if x != nil {
		return x.Resource
	}
	return nil
}

type ListOfConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  int32 `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
	Out int32 `protobuf:"varint,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *ListOfConnection) Reset() {
	*x = ListOfConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_summary_summary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfConnection) ProtoMessage() {}

func (x *ListOfConnection) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_summary_summary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfConnection.ProtoReflect.Descriptor instead.
func (*ListOfConnection) Descriptor() ([]byte, []int) {
	return file_src_proto_summary_summary_proto_rawDescGZIP(), []int{3}
}

func (x *ListOfConnection) GetIn() int32 {
	if x != nil {
		return x.In
	}
	return 0
}

func (x *ListOfConnection) GetOut() int32 {
	if x != nil {
		return x.Out
	}
	return 0
}

var File_src_proto_summary_summary_proto protoreflect.FileDescriptor

var file_src_proto_summary_summary_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x41, 0x0a, 0x0b, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xe3, 0x02,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x6c, 0x69,
	0x73, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x3b,
	0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0d, 0x6c, 0x69,
	0x73, 0x74, 0x4f, 0x66, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x33, 0x0a, 0x07, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x34, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x32, 0x45, 0x0a,
	0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x14, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x75, 0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_summary_summary_proto_rawDescOnce sync.Once
	file_src_proto_summary_summary_proto_rawDescData = file_src_proto_summary_summary_proto_rawDesc
)

func file_src_proto_summary_summary_proto_rawDescGZIP() []byte {
	file_src_proto_summary_summary_proto_rawDescOnce.Do(func() {
		file_src_proto_summary_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_summary_summary_proto_rawDescData)
	})
	return file_src_proto_summary_summary_proto_rawDescData
}

var file_src_proto_summary_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_proto_summary_summary_proto_goTypes = []interface{}{
	(*LogsRequest)(nil),      // 0: summary.LogsRequest
	(*LogsResponse)(nil),     // 1: summary.LogsResponse
	(*ListOfSource)(nil),     // 2: summary.ListOfSource
	(*ListOfConnection)(nil), // 3: summary.ListOfConnection
}
var file_src_proto_summary_summary_proto_depIdxs = []int32{
	2, // 0: summary.LogsResponse.listOfProcess:type_name -> summary.ListOfSource
	2, // 1: summary.LogsResponse.listOfFile:type_name -> summary.ListOfSource
	2, // 2: summary.LogsResponse.listOfNetwork:type_name -> summary.ListOfSource
	3, // 3: summary.LogsResponse.ingress:type_name -> summary.ListOfConnection
	3, // 4: summary.LogsResponse.egress:type_name -> summary.ListOfConnection
	0, // 5: summary.Summary.FetchLogs:input_type -> summary.LogsRequest
	1, // 6: summary.Summary.FetchLogs:output_type -> summary.LogsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_src_proto_summary_summary_proto_init() }
func file_src_proto_summary_summary_proto_init() {
	if File_src_proto_summary_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_summary_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsRequest); i {
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
		file_src_proto_summary_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponse); i {
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
		file_src_proto_summary_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfSource); i {
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
		file_src_proto_summary_summary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfConnection); i {
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
			RawDescriptor: file_src_proto_summary_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_summary_summary_proto_goTypes,
		DependencyIndexes: file_src_proto_summary_summary_proto_depIdxs,
		MessageInfos:      file_src_proto_summary_summary_proto_msgTypes,
	}.Build()
	File_src_proto_summary_summary_proto = out.File
	file_src_proto_summary_summary_proto_rawDesc = nil
	file_src_proto_summary_summary_proto_goTypes = nil
	file_src_proto_summary_summary_proto_depIdxs = nil
}
