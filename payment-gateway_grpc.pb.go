// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: payment-gateway.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentGatewayService_GetVirtualAccountInfo_FullMethodName      = "/payment.PaymentGatewayService/GetVirtualAccountInfo"
	PaymentGatewayService_GetPartneredBanks_FullMethodName          = "/payment.PaymentGatewayService/GetPartneredBanks"
	PaymentGatewayService_GetAvailablePaymentMethods_FullMethodName = "/payment.PaymentGatewayService/GetAvailablePaymentMethods"
)

// PaymentGatewayServiceClient is the client API for PaymentGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentGatewayServiceClient interface {
	GetVirtualAccountInfo(ctx context.Context, in *GetVirtualAccountInfoRequest, opts ...grpc.CallOption) (*GetVirtualAccountInfoResponse, error)
	// Helper RPC
	GetPartneredBanks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetAvailablePaymentMethods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type paymentGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentGatewayServiceClient(cc grpc.ClientConnInterface) PaymentGatewayServiceClient {
	return &paymentGatewayServiceClient{cc}
}

func (c *paymentGatewayServiceClient) GetVirtualAccountInfo(ctx context.Context, in *GetVirtualAccountInfoRequest, opts ...grpc.CallOption) (*GetVirtualAccountInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVirtualAccountInfoResponse)
	err := c.cc.Invoke(ctx, PaymentGatewayService_GetVirtualAccountInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPartneredBanks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, PaymentGatewayService_GetPartneredBanks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetAvailablePaymentMethods(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, PaymentGatewayService_GetAvailablePaymentMethods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGatewayServiceServer is the server API for PaymentGatewayService service.
// All implementations must embed UnimplementedPaymentGatewayServiceServer
// for forward compatibility.
type PaymentGatewayServiceServer interface {
	GetVirtualAccountInfo(context.Context, *GetVirtualAccountInfoRequest) (*GetVirtualAccountInfoResponse, error)
	// Helper RPC
	GetPartneredBanks(context.Context, *Empty) (*Empty, error)
	GetAvailablePaymentMethods(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedPaymentGatewayServiceServer()
}

// UnimplementedPaymentGatewayServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentGatewayServiceServer struct{}

func (UnimplementedPaymentGatewayServiceServer) GetVirtualAccountInfo(context.Context, *GetVirtualAccountInfoRequest) (*GetVirtualAccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVirtualAccountInfo not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPartneredBanks(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartneredBanks not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetAvailablePaymentMethods(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailablePaymentMethods not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) mustEmbedUnimplementedPaymentGatewayServiceServer() {}
func (UnimplementedPaymentGatewayServiceServer) testEmbeddedByValue()                               {}

// UnsafePaymentGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentGatewayServiceServer will
// result in compilation errors.
type UnsafePaymentGatewayServiceServer interface {
	mustEmbedUnimplementedPaymentGatewayServiceServer()
}

func RegisterPaymentGatewayServiceServer(s grpc.ServiceRegistrar, srv PaymentGatewayServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentGatewayServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentGatewayService_ServiceDesc, srv)
}

func _PaymentGatewayService_GetVirtualAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVirtualAccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetVirtualAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGatewayService_GetVirtualAccountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetVirtualAccountInfo(ctx, req.(*GetVirtualAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPartneredBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPartneredBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGatewayService_GetPartneredBanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPartneredBanks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetAvailablePaymentMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetAvailablePaymentMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGatewayService_GetAvailablePaymentMethods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetAvailablePaymentMethods(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentGatewayService_ServiceDesc is the grpc.ServiceDesc for PaymentGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentGatewayService",
	HandlerType: (*PaymentGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVirtualAccountInfo",
			Handler:    _PaymentGatewayService_GetVirtualAccountInfo_Handler,
		},
		{
			MethodName: "GetPartneredBanks",
			Handler:    _PaymentGatewayService_GetPartneredBanks_Handler,
		},
		{
			MethodName: "GetAvailablePaymentMethods",
			Handler:    _PaymentGatewayService_GetAvailablePaymentMethods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment-gateway.proto",
}
