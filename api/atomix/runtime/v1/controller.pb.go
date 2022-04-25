// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/controller.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Cluster struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Data    []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c82018edebe4a3, []int{0}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return m.Size()
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Cluster) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Cluster) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetClusterRequest struct {
	Store string `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c82018edebe4a3, []int{1}
}
func (m *GetClusterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetClusterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterRequest.Merge(m, src)
}
func (m *GetClusterRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterRequest proto.InternalMessageInfo

func (m *GetClusterRequest) GetStore() string {
	if m != nil {
		return m.Store
	}
	return ""
}

type GetClusterResponse struct {
	Cluster Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster"`
}

func (m *GetClusterResponse) Reset()         { *m = GetClusterResponse{} }
func (m *GetClusterResponse) String() string { return proto.CompactTextString(m) }
func (*GetClusterResponse) ProtoMessage()    {}
func (*GetClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c82018edebe4a3, []int{2}
}
func (m *GetClusterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetClusterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterResponse.Merge(m, src)
}
func (m *GetClusterResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterResponse proto.InternalMessageInfo

func (m *GetClusterResponse) GetCluster() Cluster {
	if m != nil {
		return m.Cluster
	}
	return Cluster{}
}

type ListClustersRequest struct {
	Watch bool   `protobuf:"varint,1,opt,name=watch,proto3" json:"watch,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c82018edebe4a3, []int{3}
}
func (m *ListClustersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListClustersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListClustersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListClustersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersRequest.Merge(m, src)
}
func (m *ListClustersRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListClustersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersRequest proto.InternalMessageInfo

func (m *ListClustersRequest) GetWatch() bool {
	if m != nil {
		return m.Watch
	}
	return false
}

func (m *ListClustersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListClustersResponse struct {
	Cluster Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68c82018edebe4a3, []int{4}
}
func (m *ListClustersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListClustersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersResponse.Merge(m, src)
}
func (m *ListClustersResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersResponse proto.InternalMessageInfo

func (m *ListClustersResponse) GetCluster() Cluster {
	if m != nil {
		return m.Cluster
	}
	return Cluster{}
}

func init() {
	proto.RegisterType((*Cluster)(nil), "atomix.runtime.v1.Cluster")
	proto.RegisterType((*GetClusterRequest)(nil), "atomix.runtime.v1.GetClusterRequest")
	proto.RegisterType((*GetClusterResponse)(nil), "atomix.runtime.v1.GetClusterResponse")
	proto.RegisterType((*ListClustersRequest)(nil), "atomix.runtime.v1.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "atomix.runtime.v1.ListClustersResponse")
}

func init() {
	proto.RegisterFile("atomix/runtime/v1/controller.proto", fileDescriptor_68c82018edebe4a3)
}

var fileDescriptor_68c82018edebe4a3 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x4f, 0x32, 0x41,
	0x10, 0xc6, 0x6f, 0x79, 0x79, 0x45, 0x47, 0x1a, 0x56, 0x8a, 0xcb, 0x15, 0x2b, 0xd9, 0xf8, 0x07,
	0x9b, 0x3b, 0xc1, 0xce, 0xc6, 0x04, 0x0a, 0x1b, 0x0b, 0xb3, 0x9d, 0x95, 0x59, 0x71, 0x83, 0x97,
	0xc0, 0x2d, 0xee, 0x2e, 0xa8, 0xdf, 0xc2, 0x8f, 0x45, 0xac, 0x28, 0xad, 0x8c, 0x81, 0x2f, 0x62,
	0x6e, 0xf7, 0x0e, 0x30, 0x5c, 0x42, 0x61, 0x37, 0x33, 0xf7, 0x9b, 0x79, 0xe6, 0x99, 0x5b, 0xa0,
	0xdc, 0xc8, 0x61, 0xfc, 0x1a, 0xa9, 0x71, 0x62, 0xe2, 0xa1, 0x88, 0x26, 0xad, 0xa8, 0x27, 0x13,
	0xa3, 0xe4, 0x60, 0x20, 0x54, 0x38, 0x52, 0xd2, 0x48, 0x5c, 0x73, 0x4c, 0x98, 0x31, 0xe1, 0xa4,
	0x15, 0xd4, 0xfb, 0xb2, 0x2f, 0xed, 0xd7, 0x28, 0x8d, 0x1c, 0x48, 0xef, 0xa1, 0xd2, 0x1d, 0x8c,
	0xb5, 0x11, 0x0a, 0x63, 0x28, 0x27, 0x7c, 0x28, 0x7c, 0xd4, 0x40, 0xcd, 0x3d, 0x66, 0xe3, 0xb4,
	0x66, 0xde, 0x46, 0xc2, 0x2f, 0xb9, 0x5a, 0x1a, 0x63, 0x1f, 0x2a, 0x13, 0xa1, 0x74, 0x2c, 0x13,
	0xff, 0x9f, 0x2d, 0xe7, 0x69, 0x4a, 0x3f, 0x72, 0xc3, 0xfd, 0x72, 0x03, 0x35, 0xab, 0xcc, 0xc6,
	0xf4, 0x0c, 0x6a, 0xd7, 0xc2, 0x64, 0x1a, 0x4c, 0x3c, 0x8f, 0x85, 0x36, 0xb8, 0x0e, 0xff, 0xb5,
	0x91, 0x2a, 0xd7, 0x72, 0x09, 0xbd, 0x05, 0xbc, 0x8e, 0xea, 0x91, 0x4c, 0xb4, 0xc0, 0x97, 0x50,
	0xe9, 0xb9, 0x92, 0xa5, 0xf7, 0xdb, 0x41, 0xb8, 0x61, 0x2e, 0xcc, 0x9a, 0x3a, 0xe5, 0xe9, 0xd7,
	0xa1, 0xc7, 0xf2, 0x06, 0x7a, 0x05, 0x07, 0x37, 0xb1, 0xce, 0x47, 0xea, 0x35, 0xf9, 0x17, 0x6e,
	0x7a, 0x4f, 0x76, 0xe0, 0x2e, 0x73, 0xc9, 0xd2, 0x7f, 0x69, 0xe5, 0x9f, 0x32, 0xa8, 0xff, 0x1e,
	0xf0, 0xf7, 0xa5, 0xda, 0x1f, 0x08, 0xa0, 0xbb, 0xfc, 0x61, 0xf8, 0x0e, 0x60, 0xe5, 0x1a, 0x1f,
	0x15, 0xcc, 0xd9, 0xb8, 0x5f, 0x70, 0xbc, 0x85, 0xca, 0xb6, 0xe4, 0x50, 0x5d, 0xdf, 0x1e, 0x9f,
	0x14, 0xb4, 0x15, 0xdc, 0x27, 0x38, 0xdd, 0xca, 0x39, 0x81, 0x73, 0xd4, 0xf1, 0xa7, 0x73, 0x82,
	0x66, 0x73, 0x82, 0xbe, 0xe7, 0x04, 0xbd, 0x2f, 0x88, 0x37, 0x5b, 0x10, 0xef, 0x73, 0x41, 0xbc,
	0x87, 0x1d, 0xfb, 0xc0, 0x2e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x3f, 0xb8, 0x44, 0xaf,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerClient interface {
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (Controller_ListClustersClient, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error) {
	out := new(GetClusterResponse)
	err := c.cc.Invoke(ctx, "/atomix.runtime.v1.Controller/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (Controller_ListClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Controller_serviceDesc.Streams[0], "/atomix.runtime.v1.Controller/ListClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerListClustersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Controller_ListClustersClient interface {
	Recv() (*ListClustersResponse, error)
	grpc.ClientStream
}

type controllerListClustersClient struct {
	grpc.ClientStream
}

func (x *controllerListClustersClient) Recv() (*ListClustersResponse, error) {
	m := new(ListClustersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerServer is the server API for Controller service.
type ControllerServer interface {
	GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error)
	ListClusters(*ListClustersRequest, Controller_ListClustersServer) error
}

// UnimplementedControllerServer can be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (*UnimplementedControllerServer) GetCluster(ctx context.Context, req *GetClusterRequest) (*GetClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (*UnimplementedControllerServer) ListClusters(req *ListClustersRequest, srv Controller_ListClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.runtime.v1.Controller/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_ListClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListClustersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServer).ListClusters(m, &controllerListClustersServer{stream})
}

type Controller_ListClustersServer interface {
	Send(*ListClustersResponse) error
	grpc.ServerStream
}

type controllerListClustersServer struct {
	grpc.ServerStream
}

func (x *controllerListClustersServer) Send(m *ListClustersResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCluster",
			Handler:    _Controller_GetCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListClusters",
			Handler:       _Controller_ListClusters_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "atomix/runtime/v1/controller.proto",
}

func (m *Cluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintController(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintController(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintController(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintController(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetClusterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetClusterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetClusterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Store) > 0 {
		i -= len(m.Store)
		copy(dAtA[i:], m.Store)
		i = encodeVarintController(dAtA, i, uint64(len(m.Store)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetClusterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetClusterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetClusterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Cluster.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintController(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ListClustersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListClustersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListClustersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintController(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Watch {
		i--
		if m.Watch {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListClustersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListClustersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListClustersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Cluster.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintController(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintController(dAtA []byte, offset int, v uint64) int {
	offset -= sovController(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	return n
}

func (m *GetClusterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Store)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	return n
}

func (m *GetClusterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cluster.Size()
	n += 1 + l + sovController(uint64(l))
	return n
}

func (m *ListClustersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Watch {
		n += 2
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovController(uint64(l))
	}
	return n
}

func (m *ListClustersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Cluster.Size()
	n += 1 + l + sovController(uint64(l))
	return n
}

func sovController(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozController(x uint64) (n int) {
	return sovController(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetClusterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetClusterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetClusterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Store = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetClusterResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetClusterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetClusterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cluster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListClustersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListClustersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListClustersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Watch = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListClustersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowController
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListClustersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListClustersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowController
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthController
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthController
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cluster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipController(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthController
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipController(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowController
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowController
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowController
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthController
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupController
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthController
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthController        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowController          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupController = fmt.Errorf("proto: unexpected end of group")
)
