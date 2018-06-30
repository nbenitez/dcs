// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcs.proto

package dcspb // import "github.com/Debian/dcs/internal/proto/dcspb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sourcebackendpb "github.com/Debian/dcs/internal/proto/sourcebackendpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error_ErrorType int32

const (
	Error_CANCELLED           Error_ErrorType = 0
	Error_BACKEND_UNAVAILABLE Error_ErrorType = 1
	Error_FAILED              Error_ErrorType = 2
	Error_INVALID_QUERY       Error_ErrorType = 3
)

var Error_ErrorType_name = map[int32]string{
	0: "CANCELLED",
	1: "BACKEND_UNAVAILABLE",
	2: "FAILED",
	3: "INVALID_QUERY",
}
var Error_ErrorType_value = map[string]int32{
	"CANCELLED":           0,
	"BACKEND_UNAVAILABLE": 1,
	"FAILED":              2,
	"INVALID_QUERY":       3,
}

func (x Error_ErrorType) String() string {
	return proto.EnumName(Error_ErrorType_name, int32(x))
}
func (Error_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{1, 0}
}

type Event_Type int32

const (
	Event_ERROR      Event_Type = 0
	Event_PROGRESS   Event_Type = 1
	Event_MATCH      Event_Type = 2
	Event_PAGINATION Event_Type = 3
	Event_DONE       Event_Type = 4
)

var Event_Type_name = map[int32]string{
	0: "ERROR",
	1: "PROGRESS",
	2: "MATCH",
	3: "PAGINATION",
	4: "DONE",
}
var Event_Type_value = map[string]int32{
	"ERROR":      0,
	"PROGRESS":   1,
	"MATCH":      2,
	"PAGINATION": 3,
	"DONE":       4,
}

func (x Event_Type) String() string {
	return proto.EnumName(Event_Type_name, int32(x))
}
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{4, 0}
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{0}
}
func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (dst *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(dst, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Error struct {
	Type                 Error_ErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=dcspb.Error_ErrorType" json:"type,omitempty"`
	Message              string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{1}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetType() Error_ErrorType {
	if m != nil {
		return m.Type
	}
	return Error_CANCELLED
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Progress struct {
	QueryId              string   `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	FilesProcessed       int64    `protobuf:"varint,2,opt,name=files_processed,json=filesProcessed,proto3" json:"files_processed,omitempty"`
	FilesTotal           int64    `protobuf:"varint,3,opt,name=files_total,json=filesTotal,proto3" json:"files_total,omitempty"`
	Results              int64    `protobuf:"varint,4,opt,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{2}
}
func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (dst *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(dst, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Progress) GetFilesProcessed() int64 {
	if m != nil {
		return m.FilesProcessed
	}
	return 0
}

func (m *Progress) GetFilesTotal() int64 {
	if m != nil {
		return m.FilesTotal
	}
	return 0
}

func (m *Progress) GetResults() int64 {
	if m != nil {
		return m.Results
	}
	return 0
}

type Pagination struct {
	QueryId              string   `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	ResultPages          int64    `protobuf:"varint,2,opt,name=result_pages,json=resultPages,proto3" json:"result_pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{3}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Pagination) GetResultPages() int64 {
	if m != nil {
		return m.ResultPages
	}
	return 0
}

type Event struct {
	// Types that are valid to be assigned to Data:
	//	*Event_Error
	//	*Event_Progress
	//	*Event_Match
	//	*Event_Pagination
	Data                 isEvent_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcs_56758c9931024b32, []int{4}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Data interface {
	isEvent_Data()
}

type Event_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}
type Event_Progress struct {
	Progress *Progress `protobuf:"bytes,2,opt,name=progress,proto3,oneof"`
}
type Event_Match struct {
	Match *sourcebackendpb.Match `protobuf:"bytes,3,opt,name=match,proto3,oneof"`
}
type Event_Pagination struct {
	Pagination *Pagination `protobuf:"bytes,4,opt,name=pagination,proto3,oneof"`
}

func (*Event_Error) isEvent_Data()      {}
func (*Event_Progress) isEvent_Data()   {}
func (*Event_Match) isEvent_Data()      {}
func (*Event_Pagination) isEvent_Data() {}

func (m *Event) GetData() isEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetError() *Error {
	if x, ok := m.GetData().(*Event_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Event) GetProgress() *Progress {
	if x, ok := m.GetData().(*Event_Progress); ok {
		return x.Progress
	}
	return nil
}

func (m *Event) GetMatch() *sourcebackendpb.Match {
	if x, ok := m.GetData().(*Event_Match); ok {
		return x.Match
	}
	return nil
}

func (m *Event) GetPagination() *Pagination {
	if x, ok := m.GetData().(*Event_Pagination); ok {
		return x.Pagination
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Error)(nil),
		(*Event_Progress)(nil),
		(*Event_Match)(nil),
		(*Event_Pagination)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_Error:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Event_Progress:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Progress); err != nil {
			return err
		}
	case *Event_Match:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Match); err != nil {
			return err
		}
	case *Event_Pagination:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pagination); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Data has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 1: // data.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Error{msg}
		return true, err
	case 2: // data.progress
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Progress)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Progress{msg}
		return true, err
	case 3: // data.match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(sourcebackendpb.Match)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Match{msg}
		return true, err
	case 4: // data.pagination
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pagination)
		err := b.DecodeMessage(msg)
		m.Data = &Event_Pagination{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// data
	switch x := m.Data.(type) {
	case *Event_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Progress:
		s := proto.Size(x.Progress)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Match:
		s := proto.Size(x.Match)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Pagination:
		s := proto.Size(x.Pagination)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "dcspb.SearchRequest")
	proto.RegisterType((*Error)(nil), "dcspb.Error")
	proto.RegisterType((*Progress)(nil), "dcspb.Progress")
	proto.RegisterType((*Pagination)(nil), "dcspb.Pagination")
	proto.RegisterType((*Event)(nil), "dcspb.Event")
	proto.RegisterEnum("dcspb.Error_ErrorType", Error_ErrorType_name, Error_ErrorType_value)
	proto.RegisterEnum("dcspb.Event_Type", Event_Type_name, Event_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DCSClient is the client API for DCS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCSClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DCS_SearchClient, error)
}

type dCSClient struct {
	cc *grpc.ClientConn
}

func NewDCSClient(cc *grpc.ClientConn) DCSClient {
	return &dCSClient{cc}
}

func (c *dCSClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (DCS_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DCS_serviceDesc.Streams[0], "/dcspb.DCS/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &dCSSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DCS_SearchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type dCSSearchClient struct {
	grpc.ClientStream
}

func (x *dCSSearchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DCSServer is the server API for DCS service.
type DCSServer interface {
	Search(*SearchRequest, DCS_SearchServer) error
}

func RegisterDCSServer(s *grpc.Server, srv DCSServer) {
	s.RegisterService(&_DCS_serviceDesc, srv)
}

func _DCS_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DCSServer).Search(m, &dCSSearchServer{stream})
}

type DCS_SearchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type dCSSearchServer struct {
	grpc.ServerStream
}

func (x *dCSSearchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _DCS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcspb.DCS",
	HandlerType: (*DCSServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _DCS_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dcs.proto",
}

func init() { proto.RegisterFile("dcs.proto", fileDescriptor_dcs_56758c9931024b32) }

var fileDescriptor_dcs_56758c9931024b32 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xb6, 0x13, 0x3b, 0x4d, 0x8e, 0x7b, 0x71, 0xe7, 0xaf, 0xfa, 0x87, 0x6e, 0x00, 0x03, 0x02,
	0x55, 0x60, 0x57, 0xee, 0x82, 0xb5, 0x13, 0x9b, 0xc6, 0x90, 0x3a, 0x66, 0x92, 0x56, 0x82, 0x4d,
	0x34, 0xb6, 0x87, 0xd4, 0x22, 0xb5, 0xdd, 0x99, 0x09, 0x52, 0x1f, 0x81, 0x15, 0xcf, 0xc0, 0x9b,
	0x22, 0x8f, 0x93, 0xa8, 0x65, 0xc1, 0x26, 0xd2, 0x77, 0x99, 0x73, 0xf9, 0x72, 0x0c, 0xbd, 0x2c,
	0xe5, 0x76, 0xc5, 0x4a, 0x51, 0x22, 0x3d, 0x4b, 0x79, 0x95, 0x9c, 0xbc, 0xe0, 0xe5, 0x8a, 0xa5,
	0x34, 0x21, 0xe9, 0x77, 0x5a, 0x64, 0x55, 0xe2, 0x3c, 0xc2, 0x8d, 0xd7, 0x7a, 0x05, 0x7b, 0x53,
	0x4a, 0x58, 0x7a, 0x83, 0xe9, 0xdd, 0x8a, 0x72, 0x81, 0x8e, 0x40, 0xbf, 0x5b, 0x51, 0x76, 0xdf,
	0x57, 0x9f, 0xa9, 0x6f, 0x7a, 0xb8, 0x01, 0xd6, 0x6f, 0x15, 0xf4, 0x80, 0xb1, 0x92, 0xa1, 0x53,
	0xd0, 0xc4, 0x7d, 0x45, 0xa5, 0xbc, 0xef, 0x1e, 0xdb, 0xb2, 0x97, 0x2d, 0xb5, 0xe6, 0x77, 0x76,
	0x5f, 0x51, 0x2c, 0x3d, 0xa8, 0x0f, 0x3b, 0xb7, 0x94, 0x73, 0xb2, 0xa0, 0xfd, 0x96, 0xac, 0xb6,
	0x81, 0x16, 0x86, 0xde, 0xd6, 0x8c, 0xf6, 0xa0, 0x37, 0xf4, 0xa2, 0x61, 0x30, 0x1e, 0x07, 0xbe,
	0xa9, 0xa0, 0xff, 0xe1, 0xbf, 0x81, 0x37, 0xfc, 0x14, 0x44, 0xfe, 0xfc, 0x2a, 0xf2, 0xae, 0xbd,
	0x70, 0xec, 0x0d, 0xc6, 0x81, 0xa9, 0x22, 0x80, 0xce, 0x07, 0x2f, 0xac, 0x4d, 0x2d, 0x74, 0x08,
	0x7b, 0x61, 0x74, 0xed, 0x8d, 0x43, 0x7f, 0xfe, 0xf9, 0x2a, 0xc0, 0x5f, 0xcc, 0xb6, 0xf5, 0x53,
	0x85, 0x6e, 0xcc, 0xca, 0x05, 0xa3, 0x9c, 0xa3, 0x27, 0xd0, 0x95, 0x93, 0xcf, 0xf3, 0x6c, 0xbd,
	0xc9, 0x8e, 0xc4, 0x61, 0x86, 0x5e, 0xc3, 0xc1, 0xb7, 0x7c, 0x49, 0xf9, 0xbc, 0x62, 0x65, 0x4a,
	0x39, 0xa7, 0x99, 0x9c, 0xae, 0x8d, 0xf7, 0x25, 0x1d, 0x6f, 0x58, 0xf4, 0x14, 0x8c, 0xc6, 0x28,
	0x4a, 0x41, 0x96, 0xfd, 0xb6, 0x34, 0x81, 0xa4, 0x66, 0x35, 0x53, 0xef, 0xc7, 0x28, 0x5f, 0x2d,
	0x05, 0xef, 0x6b, 0x52, 0xdc, 0x40, 0xeb, 0x23, 0x40, 0x4c, 0x16, 0x79, 0x41, 0x44, 0x5e, 0x16,
	0xff, 0x1a, 0xe6, 0x39, 0xec, 0x36, 0x6f, 0xe6, 0x15, 0x59, 0x50, 0xbe, 0x9e, 0xc4, 0x68, 0xb8,
	0xb8, 0xa6, 0xac, 0x5f, 0x2d, 0xd0, 0x83, 0x1f, 0xb4, 0x10, 0xe8, 0x25, 0xe8, 0xb4, 0x4e, 0x4d,
	0x16, 0x31, 0xdc, 0xdd, 0x87, 0xe1, 0x8f, 0x14, 0xdc, 0x88, 0xe8, 0x1d, 0x74, 0xab, 0x75, 0x0c,
	0xb2, 0x9c, 0xe1, 0x1e, 0xac, 0x8d, 0x9b, 0x74, 0x46, 0x0a, 0xde, 0x5a, 0x90, 0x0d, 0xfa, 0x2d,
	0x11, 0xe9, 0x8d, 0xdc, 0xcf, 0x70, 0x8f, 0xed, 0xbf, 0xce, 0xc6, 0xbe, 0xac, 0xd5, 0xba, 0xbc,
	0xb4, 0xa1, 0x73, 0x80, 0x6a, 0xbb, 0x9a, 0xdc, 0xdb, 0x70, 0x0f, 0x37, 0x0d, 0xb6, 0xc2, 0x48,
	0xc1, 0x0f, 0x6c, 0x96, 0x0f, 0x9a, 0xfc, 0xab, 0x7b, 0xa0, 0x07, 0x18, 0x4f, 0xb0, 0xa9, 0xa0,
	0x5d, 0xe8, 0xc6, 0x78, 0x72, 0x81, 0x83, 0xe9, 0xd4, 0x54, 0x6b, 0xe1, 0xd2, 0x9b, 0x0d, 0x47,
	0x66, 0x0b, 0xed, 0x03, 0xc4, 0xde, 0x45, 0x18, 0x79, 0xb3, 0x70, 0x12, 0x99, 0x6d, 0xd4, 0x05,
	0xcd, 0x9f, 0x44, 0x81, 0xa9, 0x0d, 0x3a, 0xa0, 0x65, 0x44, 0x10, 0xf7, 0x3d, 0xb4, 0xfd, 0xe1,
	0x14, 0x9d, 0x41, 0xa7, 0xb9, 0x5d, 0x74, 0xb4, 0xee, 0xff, 0xe8, 0x94, 0x4f, 0xb6, 0xf9, 0xd4,
	0xe1, 0x59, 0xca, 0x99, 0x3a, 0x78, 0xfb, 0xf5, 0x74, 0x91, 0x8b, 0x9b, 0x55, 0x62, 0xa7, 0xe5,
	0xad, 0xe3, 0xd3, 0x24, 0x27, 0x85, 0x93, 0xa5, 0xdc, 0xc9, 0x0b, 0x41, 0x59, 0x41, 0x96, 0x8e,
	0xfc, 0x2a, 0x1c, 0xf9, 0x2e, 0xe9, 0x48, 0x70, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x55, 0x21,
	0x55, 0x8e, 0x5b, 0x03, 0x00, 0x00,
}