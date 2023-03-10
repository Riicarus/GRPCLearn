// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: rpc/route/route.proto

package route

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CpuServiceClient is the client API for CpuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CpuServiceClient interface {
	// unary
	SearchCPU(ctx context.Context, in *CpuOfNameRequest, opts ...grpc.CallOption) (*Cpu, error)
	// server side stream
	ListCPUOfOneBrand(ctx context.Context, in *CpuOfBrandRequest, opts ...grpc.CallOption) (CpuService_ListCPUOfOneBrandClient, error)
	// user side stream
	CountNumber(ctx context.Context, opts ...grpc.CallOption) (CpuService_CountNumberClient, error)
	// bi-directional stream
	ListCPUOfNames(ctx context.Context, opts ...grpc.CallOption) (CpuService_ListCPUOfNamesClient, error)
}

type cpuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCpuServiceClient(cc grpc.ClientConnInterface) CpuServiceClient {
	return &cpuServiceClient{cc}
}

func (c *cpuServiceClient) SearchCPU(ctx context.Context, in *CpuOfNameRequest, opts ...grpc.CallOption) (*Cpu, error) {
	out := new(Cpu)
	err := c.cc.Invoke(ctx, "/route.CpuService/SearchCPU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cpuServiceClient) ListCPUOfOneBrand(ctx context.Context, in *CpuOfBrandRequest, opts ...grpc.CallOption) (CpuService_ListCPUOfOneBrandClient, error) {
	stream, err := c.cc.NewStream(ctx, &CpuService_ServiceDesc.Streams[0], "/route.CpuService/ListCPUOfOneBrand", opts...)
	if err != nil {
		return nil, err
	}
	x := &cpuServiceListCPUOfOneBrandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CpuService_ListCPUOfOneBrandClient interface {
	Recv() (*Cpu, error)
	grpc.ClientStream
}

type cpuServiceListCPUOfOneBrandClient struct {
	grpc.ClientStream
}

func (x *cpuServiceListCPUOfOneBrandClient) Recv() (*Cpu, error) {
	m := new(Cpu)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cpuServiceClient) CountNumber(ctx context.Context, opts ...grpc.CallOption) (CpuService_CountNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CpuService_ServiceDesc.Streams[1], "/route.CpuService/CountNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &cpuServiceCountNumberClient{stream}
	return x, nil
}

type CpuService_CountNumberClient interface {
	Send(*CpuOfBrandRequest) error
	CloseAndRecv() (*CpuNumberResponse, error)
	grpc.ClientStream
}

type cpuServiceCountNumberClient struct {
	grpc.ClientStream
}

func (x *cpuServiceCountNumberClient) Send(m *CpuOfBrandRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cpuServiceCountNumberClient) CloseAndRecv() (*CpuNumberResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CpuNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cpuServiceClient) ListCPUOfNames(ctx context.Context, opts ...grpc.CallOption) (CpuService_ListCPUOfNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CpuService_ServiceDesc.Streams[2], "/route.CpuService/ListCPUOfNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &cpuServiceListCPUOfNamesClient{stream}
	return x, nil
}

type CpuService_ListCPUOfNamesClient interface {
	Send(*CpuOfNameRequest) error
	Recv() (*Cpu, error)
	grpc.ClientStream
}

type cpuServiceListCPUOfNamesClient struct {
	grpc.ClientStream
}

func (x *cpuServiceListCPUOfNamesClient) Send(m *CpuOfNameRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cpuServiceListCPUOfNamesClient) Recv() (*Cpu, error) {
	m := new(Cpu)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CpuServiceServer is the server API for CpuService service.
// All implementations must embed UnimplementedCpuServiceServer
// for forward compatibility
type CpuServiceServer interface {
	// unary
	SearchCPU(context.Context, *CpuOfNameRequest) (*Cpu, error)
	// server side stream
	ListCPUOfOneBrand(*CpuOfBrandRequest, CpuService_ListCPUOfOneBrandServer) error
	// user side stream
	CountNumber(CpuService_CountNumberServer) error
	// bi-directional stream
	ListCPUOfNames(CpuService_ListCPUOfNamesServer) error
	mustEmbedUnimplementedCpuServiceServer()
}

// UnimplementedCpuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCpuServiceServer struct {
}

func (UnimplementedCpuServiceServer) SearchCPU(context.Context, *CpuOfNameRequest) (*Cpu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCPU not implemented")
}
func (UnimplementedCpuServiceServer) ListCPUOfOneBrand(*CpuOfBrandRequest, CpuService_ListCPUOfOneBrandServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCPUOfOneBrand not implemented")
}
func (UnimplementedCpuServiceServer) CountNumber(CpuService_CountNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method CountNumber not implemented")
}
func (UnimplementedCpuServiceServer) ListCPUOfNames(CpuService_ListCPUOfNamesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCPUOfNames not implemented")
}
func (UnimplementedCpuServiceServer) mustEmbedUnimplementedCpuServiceServer() {}

// UnsafeCpuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CpuServiceServer will
// result in compilation errors.
type UnsafeCpuServiceServer interface {
	mustEmbedUnimplementedCpuServiceServer()
}

func RegisterCpuServiceServer(s grpc.ServiceRegistrar, srv CpuServiceServer) {
	s.RegisterService(&CpuService_ServiceDesc, srv)
}

func _CpuService_SearchCPU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpuOfNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CpuServiceServer).SearchCPU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.CpuService/SearchCPU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CpuServiceServer).SearchCPU(ctx, req.(*CpuOfNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CpuService_ListCPUOfOneBrand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CpuOfBrandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CpuServiceServer).ListCPUOfOneBrand(m, &cpuServiceListCPUOfOneBrandServer{stream})
}

type CpuService_ListCPUOfOneBrandServer interface {
	Send(*Cpu) error
	grpc.ServerStream
}

type cpuServiceListCPUOfOneBrandServer struct {
	grpc.ServerStream
}

func (x *cpuServiceListCPUOfOneBrandServer) Send(m *Cpu) error {
	return x.ServerStream.SendMsg(m)
}

func _CpuService_CountNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CpuServiceServer).CountNumber(&cpuServiceCountNumberServer{stream})
}

type CpuService_CountNumberServer interface {
	SendAndClose(*CpuNumberResponse) error
	Recv() (*CpuOfBrandRequest, error)
	grpc.ServerStream
}

type cpuServiceCountNumberServer struct {
	grpc.ServerStream
}

func (x *cpuServiceCountNumberServer) SendAndClose(m *CpuNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cpuServiceCountNumberServer) Recv() (*CpuOfBrandRequest, error) {
	m := new(CpuOfBrandRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CpuService_ListCPUOfNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CpuServiceServer).ListCPUOfNames(&cpuServiceListCPUOfNamesServer{stream})
}

type CpuService_ListCPUOfNamesServer interface {
	Send(*Cpu) error
	Recv() (*CpuOfNameRequest, error)
	grpc.ServerStream
}

type cpuServiceListCPUOfNamesServer struct {
	grpc.ServerStream
}

func (x *cpuServiceListCPUOfNamesServer) Send(m *Cpu) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cpuServiceListCPUOfNamesServer) Recv() (*CpuOfNameRequest, error) {
	m := new(CpuOfNameRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CpuService_ServiceDesc is the grpc.ServiceDesc for CpuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CpuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.CpuService",
	HandlerType: (*CpuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchCPU",
			Handler:    _CpuService_SearchCPU_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCPUOfOneBrand",
			Handler:       _CpuService_ListCPUOfOneBrand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CountNumber",
			Handler:       _CpuService_CountNumber_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListCPUOfNames",
			Handler:       _CpuService_ListCPUOfNames_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/route/route.proto",
}
