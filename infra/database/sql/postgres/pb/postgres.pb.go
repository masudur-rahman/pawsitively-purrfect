// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: proto/database/postgres.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdParams) Reset() {
	*x = IdParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdParams) ProtoMessage() {}

func (x *IdParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdParams.ProtoReflect.Descriptor instead.
func (*IdParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{0}
}

func (x *IdParams) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *IdParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FilterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string     `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Filter *anypb.Any `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *FilterParams) Reset() {
	*x = FilterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterParams) ProtoMessage() {}

func (x *FilterParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterParams.ProtoReflect.Descriptor instead.
func (*FilterParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{1}
}

func (x *FilterParams) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *FilterParams) GetFilter() *anypb.Any {
	if x != nil {
		return x.Filter
	}
	return nil
}

type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *anypb.Any `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{2}
}

func (x *RecordResponse) GetRecord() *anypb.Any {
	if x != nil {
		return x.Record
	}
	return nil
}

type RecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*RecordResponse `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *RecordsResponse) Reset() {
	*x = RecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordsResponse) ProtoMessage() {}

func (x *RecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordsResponse.ProtoReflect.Descriptor instead.
func (*RecordsResponse) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{3}
}

func (x *RecordsResponse) GetRecords() []*RecordResponse {
	if x != nil {
		return x.Records
	}
	return nil
}

type CreateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string     `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Bean  *anypb.Any `protobuf:"bytes,2,opt,name=bean,proto3" json:"bean,omitempty"`
}

func (x *CreateParams) Reset() {
	*x = CreateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParams) ProtoMessage() {}

func (x *CreateParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParams.ProtoReflect.Descriptor instead.
func (*CreateParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{4}
}

func (x *CreateParams) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *CreateParams) GetBean() *anypb.Any {
	if x != nil {
		return x.Bean
	}
	return nil
}

type UpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string     `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Id    string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Bean  *anypb.Any `protobuf:"bytes,3,opt,name=bean,proto3" json:"bean,omitempty"`
}

func (x *UpdateParams) Reset() {
	*x = UpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateParams) ProtoMessage() {}

func (x *UpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateParams.ProtoReflect.Descriptor instead.
func (*UpdateParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateParams) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *UpdateParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateParams) GetBean() *anypb.Any {
	if x != nil {
		return x.Bean
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{6}
}

type QueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string       `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Args  []*anypb.Any `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *QueryParams) Reset() {
	*x = QueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParams) ProtoMessage() {}

func (x *QueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParams.ProtoReflect.Descriptor instead.
func (*QueryParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{7}
}

func (x *QueryParams) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryParams) GetArgs() []*anypb.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *anypb.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{8}
}

func (x *QueryResponse) GetResult() *anypb.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

type ExecParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string       `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Args  []*anypb.Any `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ExecParams) Reset() {
	*x = ExecParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecParams) ProtoMessage() {}

func (x *ExecParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecParams.ProtoReflect.Descriptor instead.
func (*ExecParams) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{9}
}

func (x *ExecParams) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ExecParams) GetArgs() []*anypb.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

type ExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *anypb.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExecResponse) Reset() {
	*x = ExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_database_postgres_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResponse) ProtoMessage() {}

func (x *ExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_database_postgres_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResponse.ProtoReflect.Descriptor instead.
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return file_proto_database_postgres_proto_rawDescGZIP(), []int{10}
}

func (x *ExecResponse) GetResult() *anypb.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_database_postgres_proto protoreflect.FileDescriptor

var file_proto_database_postgres_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x4e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x65, 0x61, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x65, 0x61,
	0x6e, 0x22, 0x5e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x65, 0x61, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x65, 0x61,
	0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x4c, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22,
	0x3c, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe6, 0x03,
	0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x67, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69,
	0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x73, 0x75, 0x64, 0x75, 0x72, 0x2d, 0x72, 0x61, 0x68,
	0x6d, 0x61, 0x6e, 0x2f, 0x70, 0x61, 0x77, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x79, 0x2d,
	0x70, 0x75, 0x72, 0x72, 0x66, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_database_postgres_proto_rawDescOnce sync.Once
	file_proto_database_postgres_proto_rawDescData = file_proto_database_postgres_proto_rawDesc
)

func file_proto_database_postgres_proto_rawDescGZIP() []byte {
	file_proto_database_postgres_proto_rawDescOnce.Do(func() {
		file_proto_database_postgres_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_database_postgres_proto_rawDescData)
	})
	return file_proto_database_postgres_proto_rawDescData
}

var file_proto_database_postgres_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_database_postgres_proto_goTypes = []interface{}{
	(*IdParams)(nil),        // 0: database.idParams
	(*FilterParams)(nil),    // 1: database.filterParams
	(*RecordResponse)(nil),  // 2: database.recordResponse
	(*RecordsResponse)(nil), // 3: database.recordsResponse
	(*CreateParams)(nil),    // 4: database.createParams
	(*UpdateParams)(nil),    // 5: database.updateParams
	(*DeleteResponse)(nil),  // 6: database.deleteResponse
	(*QueryParams)(nil),     // 7: database.queryParams
	(*QueryResponse)(nil),   // 8: database.queryResponse
	(*ExecParams)(nil),      // 9: database.execParams
	(*ExecResponse)(nil),    // 10: database.execResponse
	(*anypb.Any)(nil),       // 11: google.protobuf.Any
}
var file_proto_database_postgres_proto_depIdxs = []int32{
	11, // 0: database.filterParams.filter:type_name -> google.protobuf.Any
	11, // 1: database.recordResponse.record:type_name -> google.protobuf.Any
	2,  // 2: database.recordsResponse.records:type_name -> database.recordResponse
	11, // 3: database.createParams.bean:type_name -> google.protobuf.Any
	11, // 4: database.updateParams.bean:type_name -> google.protobuf.Any
	11, // 5: database.queryParams.args:type_name -> google.protobuf.Any
	11, // 6: database.queryResponse.result:type_name -> google.protobuf.Any
	11, // 7: database.execParams.args:type_name -> google.protobuf.Any
	11, // 8: database.execResponse.result:type_name -> google.protobuf.Any
	0,  // 9: database.Postgres.getById:input_type -> database.idParams
	1,  // 10: database.Postgres.get:input_type -> database.filterParams
	1,  // 11: database.Postgres.find:input_type -> database.filterParams
	4,  // 12: database.Postgres.create:input_type -> database.createParams
	5,  // 13: database.Postgres.update:input_type -> database.updateParams
	0,  // 14: database.Postgres.delete:input_type -> database.idParams
	7,  // 15: database.Postgres.query:input_type -> database.queryParams
	9,  // 16: database.Postgres.exec:input_type -> database.execParams
	2,  // 17: database.Postgres.getById:output_type -> database.recordResponse
	2,  // 18: database.Postgres.get:output_type -> database.recordResponse
	3,  // 19: database.Postgres.find:output_type -> database.recordsResponse
	2,  // 20: database.Postgres.create:output_type -> database.recordResponse
	2,  // 21: database.Postgres.update:output_type -> database.recordResponse
	6,  // 22: database.Postgres.delete:output_type -> database.deleteResponse
	8,  // 23: database.Postgres.query:output_type -> database.queryResponse
	10, // 24: database.Postgres.exec:output_type -> database.execResponse
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_database_postgres_proto_init() }
func file_proto_database_postgres_proto_init() {
	if File_proto_database_postgres_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_database_postgres_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdParams); i {
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
		file_proto_database_postgres_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterParams); i {
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
		file_proto_database_postgres_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
		file_proto_database_postgres_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordsResponse); i {
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
		file_proto_database_postgres_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParams); i {
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
		file_proto_database_postgres_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateParams); i {
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
		file_proto_database_postgres_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_proto_database_postgres_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParams); i {
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
		file_proto_database_postgres_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_proto_database_postgres_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecParams); i {
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
		file_proto_database_postgres_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResponse); i {
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
			RawDescriptor: file_proto_database_postgres_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_database_postgres_proto_goTypes,
		DependencyIndexes: file_proto_database_postgres_proto_depIdxs,
		MessageInfos:      file_proto_database_postgres_proto_msgTypes,
	}.Build()
	File_proto_database_postgres_proto = out.File
	file_proto_database_postgres_proto_rawDesc = nil
	file_proto_database_postgres_proto_goTypes = nil
	file_proto_database_postgres_proto_depIdxs = nil
}