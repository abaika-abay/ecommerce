// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: protos/order/order.proto

package order

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
	OrderService_CreateOrder_FullMethodName       = "/order.OrderService/CreateOrder"
	OrderService_GetOrderByID_FullMethodName      = "/order.OrderService/GetOrderByID"
	OrderService_UpdateOrderStatus_FullMethodName = "/order.OrderService/UpdateOrderStatus"
	OrderService_ListUserOrders_FullMethodName    = "/order.OrderService/ListUserOrders"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetOrderByID(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	ListUserOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderByID(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListUserOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, OrderService_ListUserOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error)
	GetOrderByID(context.Context, *GetOrderRequest) (*OrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*OrderResponse, error)
	ListUserOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderByID(context.Context, *GetOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServiceServer) ListUserOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserOrders not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderByID(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListUserOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListUserOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _OrderService_GetOrderByID_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _OrderService_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "ListUserOrders",
			Handler:    _OrderService_ListUserOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order/order.proto",
}
