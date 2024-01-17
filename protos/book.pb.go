// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: protos/book.proto

package protos

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateBookRequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId int32  `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *CreateBookRequestBody) Reset() {
	*x = CreateBookRequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequestBody) ProtoMessage() {}

func (x *CreateBookRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_protos_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequestBody.ProtoReflect.Descriptor instead.
func (*CreateBookRequestBody) Descriptor() ([]byte, []int) {
	return file_protos_book_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequestBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookRequestBody) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type CreateBookResponseBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateBookResponseBody) Reset() {
	*x = CreateBookResponseBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponseBody) ProtoMessage() {}

func (x *CreateBookResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_protos_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponseBody.ProtoReflect.Descriptor instead.
func (*CreateBookResponseBody) Descriptor() ([]byte, []int) {
	return file_protos_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookResponseBody) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateBookResponseBody) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetBookResponseBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId int32  `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *GetBookResponseBody) Reset() {
	*x = GetBookResponseBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponseBody) ProtoMessage() {}

func (x *GetBookResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_protos_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponseBody.ProtoReflect.Descriptor instead.
func (*GetBookResponseBody) Descriptor() ([]byte, []int) {
	return file_protos_book_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookResponseBody) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBookResponseBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetBookResponseBody) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type BookID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookID) Reset() {
	*x = BookID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookID) ProtoMessage() {}

func (x *BookID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookID.ProtoReflect.Descriptor instead.
func (*BookID) Descriptor() ([]byte, []int) {
	return file_protos_book_proto_rawDescGZIP(), []int{3}
}

func (x *BookID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateBookRequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AuthorId int32  `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *UpdateBookRequestBody) Reset() {
	*x = UpdateBookRequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequestBody) ProtoMessage() {}

func (x *UpdateBookRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_protos_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequestBody.ProtoReflect.Descriptor instead.
func (*UpdateBookRequestBody) Descriptor() ([]byte, []int) {
	return file_protos_book_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBookRequestBody) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBookRequestBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBookRequestBody) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

var File_protos_book_proto protoreflect.FileDescriptor

var file_protos_book_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x56, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x58, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xae, 0x02, 0x0a, 0x04, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x1c, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x67, 0x65, 0x74,
	0x12, 0x68, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1b,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x22, 0x17, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_book_proto_rawDescOnce sync.Once
	file_protos_book_proto_rawDescData = file_protos_book_proto_rawDesc
)

func file_protos_book_proto_rawDescGZIP() []byte {
	file_protos_book_proto_rawDescOnce.Do(func() {
		file_protos_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_book_proto_rawDescData)
	})
	return file_protos_book_proto_rawDescData
}

var file_protos_book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_book_proto_goTypes = []interface{}{
	(*CreateBookRequestBody)(nil),  // 0: book.CreateBookRequestBody
	(*CreateBookResponseBody)(nil), // 1: book.CreateBookResponseBody
	(*GetBookResponseBody)(nil),    // 2: book.GetBookResponseBody
	(*BookID)(nil),                 // 3: book.BookID
	(*UpdateBookRequestBody)(nil),  // 4: book.UpdateBookRequestBody
}
var file_protos_book_proto_depIdxs = []int32{
	0, // 0: book.Book.CreateBook:input_type -> book.CreateBookRequestBody
	3, // 1: book.Book.GetBook:input_type -> book.BookID
	4, // 2: book.Book.UpdateBook:input_type -> book.UpdateBookRequestBody
	1, // 3: book.Book.CreateBook:output_type -> book.CreateBookResponseBody
	2, // 4: book.Book.GetBook:output_type -> book.GetBookResponseBody
	1, // 5: book.Book.UpdateBook:output_type -> book.CreateBookResponseBody
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_book_proto_init() }
func file_protos_book_proto_init() {
	if File_protos_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequestBody); i {
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
		file_protos_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResponseBody); i {
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
		file_protos_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResponseBody); i {
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
		file_protos_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookID); i {
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
		file_protos_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequestBody); i {
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
			RawDescriptor: file_protos_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_book_proto_goTypes,
		DependencyIndexes: file_protos_book_proto_depIdxs,
		MessageInfos:      file_protos_book_proto_msgTypes,
	}.Build()
	File_protos_book_proto = out.File
	file_protos_book_proto_rawDesc = nil
	file_protos_book_proto_goTypes = nil
	file_protos_book_proto_depIdxs = nil
}
