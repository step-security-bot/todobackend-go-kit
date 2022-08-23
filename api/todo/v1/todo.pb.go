// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: todo/v1/todo.proto

package todo

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

// TodoItem is a note describing a task to be done.
type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Completed bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	Order     int32  `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_todo_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_todo_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoItem) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *TodoItem) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

var File_todo_v1_todo_proto protoreflect.FileDescriptor

var file_todo_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a,
	0x08, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x72, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x67, 0x69,
	0x6b, 0x61, 0x7a, 0x61, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x67, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x6f, 0x64, 0x6f, 0xa2, 0x02, 0x03,
	0x54, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x54, 0x6f, 0x64, 0x6f, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_v1_todo_proto_rawDescOnce sync.Once
	file_todo_v1_todo_proto_rawDescData = file_todo_v1_todo_proto_rawDesc
)

func file_todo_v1_todo_proto_rawDescGZIP() []byte {
	file_todo_v1_todo_proto_rawDescOnce.Do(func() {
		file_todo_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_v1_todo_proto_rawDescData)
	})
	return file_todo_v1_todo_proto_rawDescData
}

var file_todo_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_todo_v1_todo_proto_goTypes = []interface{}{
	(*TodoItem)(nil), // 0: todo.v1.TodoItem
}
var file_todo_v1_todo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_todo_v1_todo_proto_init() }
func file_todo_v1_todo_proto_init() {
	if File_todo_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
			RawDescriptor: file_todo_v1_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_todo_v1_todo_proto_goTypes,
		DependencyIndexes: file_todo_v1_todo_proto_depIdxs,
		MessageInfos:      file_todo_v1_todo_proto_msgTypes,
	}.Build()
	File_todo_v1_todo_proto = out.File
	file_todo_v1_todo_proto_rawDesc = nil
	file_todo_v1_todo_proto_goTypes = nil
	file_todo_v1_todo_proto_depIdxs = nil
}
