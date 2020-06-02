// Code generated by protoc-gen-kit. DO NOT EDIT.

package todo

import (
	context "context"
)

// TodoListServiceHandler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
//
// This interface is based on github.com/go-kit/kit/transport/grpc.Handler.
type TodoListServiceHandler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

// TodoListServiceKitServer is the Go kit server implementation for TodoListService service.
type TodoListServiceKitServer struct {
	*UnimplementedTodoListServiceServer

	AddItemHandler     TodoListServiceHandler
	ListItemsHandler   TodoListServiceHandler
	DeleteItemsHandler TodoListServiceHandler
	GetItemHandler     TodoListServiceHandler
	UpdateItemHandler  TodoListServiceHandler
	DeleteItemHandler  TodoListServiceHandler
}

// AddItem adds a new item to the list.
func (s TodoListServiceKitServer) AddItem(ctx context.Context, req *AddItemRequest) (*AddItemResponse, error) {
	_, resp, err := s.AddItemHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*AddItemResponse), nil
}

// ListItems returns a list of items.
func (s TodoListServiceKitServer) ListItems(ctx context.Context, req *ListItemsRequest) (*ListItemsResponse, error) {
	_, resp, err := s.ListItemsHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*ListItemsResponse), nil
}

// DeleteItems deletes all items from the list.
func (s TodoListServiceKitServer) DeleteItems(ctx context.Context, req *DeleteItemsRequest) (*DeleteItemsResponse, error) {
	_, resp, err := s.DeleteItemsHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*DeleteItemsResponse), nil
}

// GetItem returns the details of an item.
func (s TodoListServiceKitServer) GetItem(ctx context.Context, req *GetItemRequest) (*GetItemResponse, error) {
	_, resp, err := s.GetItemHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*GetItemResponse), nil
}

// UpdateItem updates an existing item.
func (s TodoListServiceKitServer) UpdateItem(ctx context.Context, req *UpdateItemRequest) (*UpdateItemResponse, error) {
	_, resp, err := s.UpdateItemHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*UpdateItemResponse), nil
}

// DeleteItem deletes an item from the list.
func (s TodoListServiceKitServer) DeleteItem(ctx context.Context, req *DeleteItemRequest) (*DeleteItemResponse, error) {
	_, resp, err := s.DeleteItemHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*DeleteItemResponse), nil
}
