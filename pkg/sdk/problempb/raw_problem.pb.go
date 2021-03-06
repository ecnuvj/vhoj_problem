// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: problempb/raw_problem.proto

package problempb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type RawProblem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawProblemId    uint64                 `protobuf:"varint,1,opt,name=raw_problem_id,json=rawProblemId,proto3" json:"raw_problem_id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	SampleInput     string                 `protobuf:"bytes,4,opt,name=sample_input,json=sampleInput,proto3" json:"sample_input,omitempty"`
	SampleOutput    string                 `protobuf:"bytes,5,opt,name=sample_output,json=sampleOutput,proto3" json:"sample_output,omitempty"`
	Input           string                 `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	Output          string                 `protobuf:"bytes,7,opt,name=output,proto3" json:"output,omitempty"`
	Hint            string                 `protobuf:"bytes,8,opt,name=hint,proto3" json:"hint,omitempty"`
	RemoteOj        string                 `protobuf:"bytes,9,opt,name=remote_oj,json=remoteOj,proto3" json:"remote_oj,omitempty"`
	RemoteProblemId string                 `protobuf:"bytes,10,opt,name=remote_problem_id,json=remoteProblemId,proto3" json:"remote_problem_id,omitempty"`
	RemoteSubmitId  string                 `protobuf:"bytes,11,opt,name=remote_submit_id,json=remoteSubmitId,proto3" json:"remote_submit_id,omitempty"`
	TimeLimit       string                 `protobuf:"bytes,12,opt,name=time_limit,json=timeLimit,proto3" json:"time_limit,omitempty"`
	MemoryLimit     string                 `protobuf:"bytes,13,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	Spj             string                 `protobuf:"bytes,14,opt,name=spj,proto3" json:"spj,omitempty"`
	Std             string                 `protobuf:"bytes,15,opt,name=std,proto3" json:"std,omitempty"`
	Source          string                 `protobuf:"bytes,16,opt,name=source,proto3" json:"source,omitempty"`
	GroupId         uint64                 `protobuf:"varint,17,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ProblemId       uint64                 `protobuf:"varint,19,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *RawProblem) Reset() {
	*x = RawProblem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problempb_raw_problem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawProblem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawProblem) ProtoMessage() {}

func (x *RawProblem) ProtoReflect() protoreflect.Message {
	mi := &file_problempb_raw_problem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawProblem.ProtoReflect.Descriptor instead.
func (*RawProblem) Descriptor() ([]byte, []int) {
	return file_problempb_raw_problem_proto_rawDescGZIP(), []int{0}
}

func (x *RawProblem) GetRawProblemId() uint64 {
	if x != nil {
		return x.RawProblemId
	}
	return 0
}

func (x *RawProblem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RawProblem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RawProblem) GetSampleInput() string {
	if x != nil {
		return x.SampleInput
	}
	return ""
}

func (x *RawProblem) GetSampleOutput() string {
	if x != nil {
		return x.SampleOutput
	}
	return ""
}

func (x *RawProblem) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *RawProblem) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *RawProblem) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *RawProblem) GetRemoteOj() string {
	if x != nil {
		return x.RemoteOj
	}
	return ""
}

func (x *RawProblem) GetRemoteProblemId() string {
	if x != nil {
		return x.RemoteProblemId
	}
	return ""
}

func (x *RawProblem) GetRemoteSubmitId() string {
	if x != nil {
		return x.RemoteSubmitId
	}
	return ""
}

func (x *RawProblem) GetTimeLimit() string {
	if x != nil {
		return x.TimeLimit
	}
	return ""
}

func (x *RawProblem) GetMemoryLimit() string {
	if x != nil {
		return x.MemoryLimit
	}
	return ""
}

func (x *RawProblem) GetSpj() string {
	if x != nil {
		return x.Spj
	}
	return ""
}

func (x *RawProblem) GetStd() string {
	if x != nil {
		return x.Std
	}
	return ""
}

func (x *RawProblem) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *RawProblem) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *RawProblem) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *RawProblem) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

var File_problempb_raw_problem_proto protoreflect.FileDescriptor

var file_problempb_raw_problem_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x77, 0x5f,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73,
	0x64, 0x6b, 0x1a, 0x27, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x04, 0x0a, 0x0a,
	0x52, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61,
	0x77, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6f, 0x6a, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4f, 0x6a, 0x12,
	0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x70, 0x6a, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x70, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x6e, 0x75, 0x76, 0x6a, 0x2f, 0x76, 0x68,
	0x6f, 0x6a, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_problempb_raw_problem_proto_rawDescOnce sync.Once
	file_problempb_raw_problem_proto_rawDescData = file_problempb_raw_problem_proto_rawDesc
)

func file_problempb_raw_problem_proto_rawDescGZIP() []byte {
	file_problempb_raw_problem_proto_rawDescOnce.Do(func() {
		file_problempb_raw_problem_proto_rawDescData = protoimpl.X.CompressGZIP(file_problempb_raw_problem_proto_rawDescData)
	})
	return file_problempb_raw_problem_proto_rawDescData
}

var file_problempb_raw_problem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_problempb_raw_problem_proto_goTypes = []interface{}{
	(*RawProblem)(nil),            // 0: sdk.RawProblem
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_problempb_raw_problem_proto_depIdxs = []int32{
	1, // 0: sdk.RawProblem.updated_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_problempb_raw_problem_proto_init() }
func file_problempb_raw_problem_proto_init() {
	if File_problempb_raw_problem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_problempb_raw_problem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawProblem); i {
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
			RawDescriptor: file_problempb_raw_problem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_problempb_raw_problem_proto_goTypes,
		DependencyIndexes: file_problempb_raw_problem_proto_depIdxs,
		MessageInfos:      file_problempb_raw_problem_proto_msgTypes,
	}.Build()
	File_problempb_raw_problem_proto = out.File
	file_problempb_raw_problem_proto_rawDesc = nil
	file_problempb_raw_problem_proto_goTypes = nil
	file_problempb_raw_problem_proto_depIdxs = nil
}
