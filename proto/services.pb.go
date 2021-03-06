// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Port struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocks              []string  `protobuf:"bytes,9,rep,name=unlocks,proto3" json:"unlocks,omitempty"`
	Code                 string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []float32 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocks() []string {
	if m != nil {
		return m.Unlocks
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GetPortsRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPortsRequest) Reset()         { *m = GetPortsRequest{} }
func (m *GetPortsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPortsRequest) ProtoMessage()    {}
func (*GetPortsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{1}
}

func (m *GetPortsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPortsRequest.Unmarshal(m, b)
}
func (m *GetPortsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPortsRequest.Marshal(b, m, deterministic)
}
func (m *GetPortsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPortsRequest.Merge(m, src)
}
func (m *GetPortsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPortsRequest.Size(m)
}
func (m *GetPortsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPortsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPortsRequest proto.InternalMessageInfo

func (m *GetPortsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetPortsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetPortsResponse struct {
	Ports                map[string]*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetPortsResponse) Reset()         { *m = GetPortsResponse{} }
func (m *GetPortsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPortsResponse) ProtoMessage()    {}
func (*GetPortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{2}
}

func (m *GetPortsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPortsResponse.Unmarshal(m, b)
}
func (m *GetPortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPortsResponse.Marshal(b, m, deterministic)
}
func (m *GetPortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPortsResponse.Merge(m, src)
}
func (m *GetPortsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPortsResponse.Size(m)
}
func (m *GetPortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPortsResponse proto.InternalMessageInfo

func (m *GetPortsResponse) GetPorts() map[string]*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PutPortsRequest struct {
	Ports                map[string]*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PutPortsRequest) Reset()         { *m = PutPortsRequest{} }
func (m *PutPortsRequest) String() string { return proto.CompactTextString(m) }
func (*PutPortsRequest) ProtoMessage()    {}
func (*PutPortsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{3}
}

func (m *PutPortsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutPortsRequest.Unmarshal(m, b)
}
func (m *PutPortsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutPortsRequest.Marshal(b, m, deterministic)
}
func (m *PutPortsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutPortsRequest.Merge(m, src)
}
func (m *PutPortsRequest) XXX_Size() int {
	return xxx_messageInfo_PutPortsRequest.Size(m)
}
func (m *PutPortsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutPortsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutPortsRequest proto.InternalMessageInfo

func (m *PutPortsRequest) GetPorts() map[string]*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PutPortsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutPortsResponse) Reset()         { *m = PutPortsResponse{} }
func (m *PutPortsResponse) String() string { return proto.CompactTextString(m) }
func (*PutPortsResponse) ProtoMessage()    {}
func (*PutPortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{4}
}

func (m *PutPortsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutPortsResponse.Unmarshal(m, b)
}
func (m *PutPortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutPortsResponse.Marshal(b, m, deterministic)
}
func (m *PutPortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutPortsResponse.Merge(m, src)
}
func (m *PutPortsResponse) XXX_Size() int {
	return xxx_messageInfo_PutPortsResponse.Size(m)
}
func (m *PutPortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutPortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutPortsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Port)(nil), "Port")
	proto.RegisterType((*GetPortsRequest)(nil), "GetPortsRequest")
	proto.RegisterType((*GetPortsResponse)(nil), "GetPortsResponse")
	proto.RegisterMapType((map[string]*Port)(nil), "GetPortsResponse.PortsEntry")
	proto.RegisterType((*PutPortsRequest)(nil), "PutPortsRequest")
	proto.RegisterMapType((map[string]*Port)(nil), "PutPortsRequest.PortsEntry")
	proto.RegisterType((*PutPortsResponse)(nil), "PutPortsResponse")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x41, 0xcf, 0x93, 0x40,
	0x10, 0x0d, 0x50, 0xfa, 0xb5, 0xd3, 0xe8, 0x47, 0x37, 0xc6, 0x6c, 0xa8, 0x07, 0xc2, 0x89, 0x13,
	0x46, 0xbc, 0x18, 0x2f, 0x35, 0x51, 0xe3, 0xb5, 0x59, 0x7f, 0x01, 0xd2, 0xa9, 0xd9, 0x14, 0x76,
	0x91, 0x5d, 0x48, 0xea, 0xc9, 0x8b, 0xf1, 0x6f, 0x9b, 0xdd, 0x05, 0xdb, 0xe2, 0xc9, 0x83, 0xb7,
	0x79, 0x6f, 0x27, 0xf3, 0xe6, 0xbd, 0x01, 0x78, 0xaa, 0xb0, 0x1b, 0x78, 0x85, 0x2a, 0x6f, 0x3b,
	0xa9, 0x65, 0xfa, 0xc3, 0x87, 0xc5, 0x41, 0x76, 0x9a, 0x10, 0x58, 0x88, 0xb2, 0x41, 0xea, 0x25,
	0x5e, 0xb6, 0x66, 0xb6, 0x36, 0x5c, 0xc5, 0xf5, 0x85, 0xfa, 0x8e, 0x33, 0x35, 0xa1, 0xf0, 0x50,
	0xc9, 0x5e, 0xe8, 0xee, 0x42, 0x03, 0x4b, 0x4f, 0x90, 0x3c, 0x83, 0xb0, 0xac, 0x79, 0xa9, 0xe8,
	0x22, 0x09, 0xb2, 0x35, 0x73, 0xc0, 0xf4, 0x77, 0xf8, 0x95, 0x4b, 0xa1, 0x68, 0x68, 0xf9, 0x09,
	0x92, 0x04, 0x36, 0x95, 0x94, 0xdd, 0x91, 0x8b, 0x52, 0xa3, 0xa2, 0xcb, 0x24, 0xc8, 0x7c, 0x76,
	0x4b, 0x91, 0x18, 0x56, 0x6d, 0x27, 0x07, 0x2e, 0x2a, 0xa4, 0x0f, 0x56, 0xec, 0x0f, 0x36, 0x6f,
	0x9a, 0x37, 0xf8, 0x5d, 0x0a, 0xa4, 0x2b, 0xf7, 0x36, 0x61, 0xa3, 0xd9, 0x8b, 0x5a, 0x56, 0x67,
	0x45, 0xd7, 0x4e, 0x73, 0x84, 0xd6, 0x91, 0x3c, 0x22, 0x85, 0xd1, 0x91, 0x3c, 0x62, 0xba, 0x87,
	0xc7, 0x4f, 0xa8, 0x4d, 0x08, 0x8a, 0xe1, 0xb7, 0x1e, 0x95, 0x26, 0xcf, 0x61, 0x29, 0x4f, 0x27,
	0x85, 0xda, 0xc6, 0x11, 0xb0, 0x11, 0x19, 0x8b, 0x35, 0x6f, 0xb8, 0xb6, 0x89, 0x04, 0xcc, 0x81,
	0xf4, 0x97, 0x07, 0xd1, 0x75, 0x82, 0x6a, 0xa5, 0x50, 0x48, 0x0a, 0x08, 0x5b, 0x43, 0x50, 0x2f,
	0x09, 0xb2, 0x4d, 0xf1, 0x22, 0x9f, 0x77, 0xe4, 0x16, 0x7d, 0x34, 0xd1, 0x31, 0xd7, 0x1a, 0xef,
	0x01, 0xae, 0x24, 0x89, 0x20, 0x38, 0xe3, 0x65, 0x3c, 0x88, 0x29, 0xc9, 0x0e, 0xc2, 0xa1, 0xac,
	0x7b, 0xb4, 0xf2, 0x9b, 0x22, 0xb4, 0x23, 0x98, 0xe3, 0xde, 0xfa, 0x6f, 0xbc, 0xf4, 0xa7, 0x07,
	0x8f, 0x87, 0xfe, 0xde, 0xcb, 0xab, 0xfb, 0x45, 0x76, 0xf9, 0xac, 0xe1, 0x7f, 0xec, 0x41, 0x20,
	0xba, 0xaa, 0x38, 0xbb, 0xc5, 0x3b, 0x78, 0xf2, 0xbe, 0xe6, 0x28, 0xf4, 0x67, 0xf7, 0x05, 0x92,
	0x97, 0xb0, 0x9a, 0x32, 0x21, 0x51, 0x3e, 0x3b, 0x41, 0xbc, 0xfd, 0x2b, 0xb0, 0x62, 0x00, 0x62,
	0x89, 0x0f, 0xb2, 0x29, 0xb9, 0xb8, 0x19, 0x33, 0x69, 0x91, 0x68, 0x6e, 0x2e, 0xde, 0xe6, 0xf3,
	0x45, 0xfe, 0x59, 0xf7, 0xcb, 0xd2, 0xfe, 0x2a, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xd0, 0x46, 0xb6, 0x3c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*GetPortsResponse, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*GetPortsResponse, error) {
	out := new(GetPortsResponse)
	err := c.cc.Invoke(ctx, "/ClientService/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	GetPorts(context.Context, *GetPortsRequest) (*GetPortsResponse, error)
}

// UnimplementedClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (*UnimplementedClientServiceServer) GetPorts(ctx context.Context, req *GetPortsRequest) (*GetPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientService/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetPorts(ctx, req.(*GetPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPorts",
			Handler:    _ClientService_GetPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

// PortsDomainServiceClient is the client API for PortsDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortsDomainServiceClient interface {
	PutPorts(ctx context.Context, in *PutPortsRequest, opts ...grpc.CallOption) (*PutPortsResponse, error)
	GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*GetPortsResponse, error)
}

type portsDomainServiceClient struct {
	cc *grpc.ClientConn
}

func NewPortsDomainServiceClient(cc *grpc.ClientConn) PortsDomainServiceClient {
	return &portsDomainServiceClient{cc}
}

func (c *portsDomainServiceClient) PutPorts(ctx context.Context, in *PutPortsRequest, opts ...grpc.CallOption) (*PutPortsResponse, error) {
	out := new(PutPortsResponse)
	err := c.cc.Invoke(ctx, "/PortsDomainService/PutPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsDomainServiceClient) GetPorts(ctx context.Context, in *GetPortsRequest, opts ...grpc.CallOption) (*GetPortsResponse, error) {
	out := new(GetPortsResponse)
	err := c.cc.Invoke(ctx, "/PortsDomainService/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsDomainServiceServer is the server API for PortsDomainService service.
type PortsDomainServiceServer interface {
	PutPorts(context.Context, *PutPortsRequest) (*PutPortsResponse, error)
	GetPorts(context.Context, *GetPortsRequest) (*GetPortsResponse, error)
}

// UnimplementedPortsDomainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPortsDomainServiceServer struct {
}

func (*UnimplementedPortsDomainServiceServer) PutPorts(ctx context.Context, req *PutPortsRequest) (*PutPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPorts not implemented")
}
func (*UnimplementedPortsDomainServiceServer) GetPorts(ctx context.Context, req *GetPortsRequest) (*GetPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}

func RegisterPortsDomainServiceServer(s *grpc.Server, srv PortsDomainServiceServer) {
	s.RegisterService(&_PortsDomainService_serviceDesc, srv)
}

func _PortsDomainService_PutPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDomainServiceServer).PutPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortsDomainService/PutPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDomainServiceServer).PutPorts(ctx, req.(*PutPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortsDomainService_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDomainServiceServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortsDomainService/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDomainServiceServer).GetPorts(ctx, req.(*GetPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortsDomainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PortsDomainService",
	HandlerType: (*PortsDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutPorts",
			Handler:    _PortsDomainService_PutPorts_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _PortsDomainService_GetPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
