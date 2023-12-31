// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: library.proto

package library

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAuthorsByBookTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetAuthorsByBookTitleRequest) Reset() {
	*x = GetAuthorsByBookTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBookTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBookTitleRequest) ProtoMessage() {}

func (x *GetAuthorsByBookTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBookTitleRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBookTitleRequest) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuthorsByBookTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetAuthorsByBookTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *GetAuthorsByBookTitleResponse) Reset() {
	*x = GetAuthorsByBookTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBookTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBookTitleResponse) ProtoMessage() {}

func (x *GetAuthorsByBookTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBookTitleResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBookTitleResponse) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthorsByBookTitleResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type GetBooksByAuthorNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetBooksByAuthorNameRequest) Reset() {
	*x = GetBooksByAuthorNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorNameRequest) ProtoMessage() {}

func (x *GetBooksByAuthorNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorNameRequest.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorNameRequest) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{4}
}

func (x *GetBooksByAuthorNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBooksByAuthorNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksByAuthorNameResponse) Reset() {
	*x = GetBooksByAuthorNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorNameResponse) ProtoMessage() {}

func (x *GetBooksByAuthorNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorNameResponse.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorNameResponse) Descriptor() ([]byte, []int) {
	return file_library_proto_rawDescGZIP(), []int{5}
}

func (x *GetBooksByAuthorNameResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_library_proto protoreflect.FileDescriptor

var file_library_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1c, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x1c, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x42, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42,
	0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x31, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0xb6, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42,
	0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65,
	0x72, 0x6e, 0x79, 0x73, 0x68, 0x65, 0x76, 0x76, 0x6c, 0x61, 0x64, 0x69, 0x73, 0x6c, 0x61, 0x76,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x2e, 0x62, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_proto_rawDescOnce sync.Once
	file_library_proto_rawDescData = file_library_proto_rawDesc
)

func file_library_proto_rawDescGZIP() []byte {
	file_library_proto_rawDescOnce.Do(func() {
		file_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_proto_rawDescData)
	})
	return file_library_proto_rawDescData
}

var file_library_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_library_proto_goTypes = []interface{}{
	(*Book)(nil),                          // 0: Book
	(*Author)(nil),                        // 1: Author
	(*GetAuthorsByBookTitleRequest)(nil),  // 2: GetAuthorsByBookTitleRequest
	(*GetAuthorsByBookTitleResponse)(nil), // 3: GetAuthorsByBookTitleResponse
	(*GetBooksByAuthorNameRequest)(nil),   // 4: GetBooksByAuthorNameRequest
	(*GetBooksByAuthorNameResponse)(nil),  // 5: GetBooksByAuthorNameResponse
}
var file_library_proto_depIdxs = []int32{
	1, // 0: GetAuthorsByBookTitleResponse.authors:type_name -> Author
	0, // 1: GetBooksByAuthorNameResponse.books:type_name -> Book
	2, // 2: Library.GetAuthorsByBookTitle:input_type -> GetAuthorsByBookTitleRequest
	4, // 3: Library.GetBooksByAuthorName:input_type -> GetBooksByAuthorNameRequest
	3, // 4: Library.GetAuthorsByBookTitle:output_type -> GetAuthorsByBookTitleResponse
	5, // 5: Library.GetBooksByAuthorName:output_type -> GetBooksByAuthorNameResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_library_proto_init() }
func file_library_proto_init() {
	if File_library_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_library_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_library_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBookTitleRequest); i {
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
		file_library_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBookTitleResponse); i {
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
		file_library_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorNameRequest); i {
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
		file_library_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorNameResponse); i {
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
			RawDescriptor: file_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_library_proto_goTypes,
		DependencyIndexes: file_library_proto_depIdxs,
		MessageInfos:      file_library_proto_msgTypes,
	}.Build()
	File_library_proto = out.File
	file_library_proto_rawDesc = nil
	file_library_proto_goTypes = nil
	file_library_proto_depIdxs = nil
}
