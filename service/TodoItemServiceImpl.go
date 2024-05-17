package service

import (
	"github.com/jguady/goRPCTutorial/proto"
)

type TodoServiceImpl struct {
	// type embedded to comply with Google lib
	proto.UnimplementedTodoServiceServer
}

func NewTodoItemServiceImpl() *TodoServiceImpl {
	return &TodoServiceImpl{}
}

// func (todoService *TodoServiceImpl) GetTodoItem(context.Context, *proto.GetTodoItemsRequest) (*proto.GetTodoItemsResponse, error) {

// }
// func (todoService *TodoServiceImpl) CreateTodoItem(context.Context, *proto.CreateTodoItemRequest) (*proto.CreateTodoItemResponse, error) {

// }
// func (todoService *TodoServiceImpl) UpdateTodoItem(context.Context, *proto.UpdateTodoItemRequest) (*proto.UpdateTodoItemResponse, error) {

// }
// func (todoService *TodoServiceImpl) DeleteTodoItem(context.Context, *proto.DeleteTodoItemRequest) (*proto.DeleteTodoItemResponse, error) {

// }
