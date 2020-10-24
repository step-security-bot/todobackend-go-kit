// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TodoListServiceClient is the client API for TodoListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoListServiceClient interface {
	// AddItem adds a new item to the list.
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	// ListItems returns a list of items.
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error)
	// DeleteItems deletes all items from the list.
	DeleteItems(ctx context.Context, in *DeleteItemsRequest, opts ...grpc.CallOption) (*DeleteItemsResponse, error)
	// GetItem returns the details of an item.
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	// UpdateItem updates an existing item.
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	// DeleteItem deletes an item from the list.
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
}

type todoListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListServiceClient(cc grpc.ClientConnInterface) TodoListServiceClient {
	return &todoListServiceClient{cc}
}

func (c *todoListServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error) {
	out := new(ListItemsResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) DeleteItems(ctx context.Context, in *DeleteItemsRequest, opts ...grpc.CallOption) (*DeleteItemsResponse, error) {
	out := new(DeleteItemsResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/DeleteItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoListService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServiceServer is the server API for TodoListService service.
// All implementations must embed UnimplementedTodoListServiceServer
// for forward compatibility
type TodoListServiceServer interface {
	// AddItem adds a new item to the list.
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	// ListItems returns a list of items.
	ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error)
	// DeleteItems deletes all items from the list.
	DeleteItems(context.Context, *DeleteItemsRequest) (*DeleteItemsResponse, error)
	// GetItem returns the details of an item.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	// UpdateItem updates an existing item.
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	// DeleteItem deletes an item from the list.
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	mustEmbedUnimplementedTodoListServiceServer()
}

// UnimplementedTodoListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoListServiceServer struct {
}

func (UnimplementedTodoListServiceServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedTodoListServiceServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedTodoListServiceServer) DeleteItems(context.Context, *DeleteItemsRequest) (*DeleteItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItems not implemented")
}
func (UnimplementedTodoListServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedTodoListServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedTodoListServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedTodoListServiceServer) mustEmbedUnimplementedTodoListServiceServer() {}

// UnsafeTodoListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoListServiceServer will
// result in compilation errors.
type UnsafeTodoListServiceServer interface {
	mustEmbedUnimplementedTodoListServiceServer()
}

func RegisterTodoListServiceServer(s grpc.ServiceRegistrar, srv TodoListServiceServer) {
	s.RegisterService(&_TodoListService_serviceDesc, srv)
}

func _TodoListService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_DeleteItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).DeleteItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/DeleteItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).DeleteItems(ctx, req.(*DeleteItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoListService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1.TodoListService",
	HandlerType: (*TodoListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _TodoListService_AddItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _TodoListService_ListItems_Handler,
		},
		{
			MethodName: "DeleteItems",
			Handler:    _TodoListService_DeleteItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _TodoListService_GetItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _TodoListService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _TodoListService_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/v1/todo_list.proto",
}
