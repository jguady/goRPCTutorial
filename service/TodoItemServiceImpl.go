package service

import (
	"context"
	"fmt"

	"github.com/jguady/goRPCTutorial/proto"
)

type TodoServiceServer struct {
	// type embedded to comply with Google lib
	proto.UnimplementedTodoServiceServer
	data map[int32]*proto.TodoItem
	id   int
}

func NewTodoItemServiceImpl() *TodoServiceServer {
	var todoServiceServer TodoServiceServer
	todoServiceServer.loadData()
	todoServiceServer.id = len(todoServiceServer.data)
	// fmt.Println("Inside New Func, Printing Data")
	// fmt.Println(data)
	fmt.Printf("The Data inside data : %v", todoServiceServer.data)
	println()
	return &todoServiceServer
}

func (todoService *TodoServiceServer) ListTodoItems(listFilter *proto.ListTodoItemFilterRequest, stream proto.TodoService_ListTodoItemsServer) error {
	fmt.Println("Listing TodoItems")
	for _, todoItem := range todoService.data {
		fmt.Printf("TodoItem: %v\n", todoItem)
		if err := stream.Send(todoItem); err != nil {
			return err
		}
	}
	return nil
}

func (todoService *TodoServiceServer) GetTodoItem(ctx context.Context, getItemRequest *proto.GetTodoItemRequest) (*proto.TodoItem, error) {
	fmt.Printf("Id is %d \n", getItemRequest.Id)
	fmt.Printf("Lookup %v \n", todoService.data)
	fmt.Printf("Lookup with # %#v \n", todoService.data)
	if _, exists := todoService.data[getItemRequest.Id]; exists {
		fmt.Println("Key Exists")
		return todoService.data[getItemRequest.Id], nil
	} else {
		return nil, fmt.Errorf("key %d does not exist in the Map", getItemRequest.Id)
	}
}
func (todoService *TodoServiceServer) CreateTodoItem(ctx context.Context, createItem *proto.CreateTodoItemRequest) (*proto.TodoItem, error) {
	todoService.id += 1
	keyId := int32(todoService.id)

	item := proto.TodoItem{Id: keyId,
		Name:     createItem.Name,
		Desc:     createItem.GetDesc(),
		DaysLeft: createItem.DaysLeft,
	}

	fmt.Printf("The create item: %v\n", createItem)
	fmt.Printf("Data before create: %v\n", todoService.data)

	todoService.data[keyId] = &item
	fmt.Printf("Data after create: %v\n", todoService.data)
	return &item, nil
}

func (todoService *TodoServiceServer) UpdateTodoItem(ctx context.Context, updateRequest *proto.UpdateTodoItemRequest) (*proto.TodoItem, error) {
	// get the item from the map
	if _, exists := todoService.data[updateRequest.Id]; exists {
		todoService.data[updateRequest.Id] = updateRequest.Item
		fmt.Println("Map Updated")
		return todoService.data[updateRequest.Id], nil
	} else {
		return nil, fmt.Errorf("key %d does not exist in the map", updateRequest.Id)
	}

}

func (todoService *TodoServiceServer) DeleteTodoItem(ctx context.Context, deleteRequest *proto.DeleteTodoItemRequest) (*proto.DeleteTodoItemResponse, error) {
	//This needs to be corrected
	delete(todoService.data, deleteRequest.Item.Id)
	// panic("TestError")
	return &proto.DeleteTodoItemResponse{}, nil
}

func (todoServiceServer *TodoServiceServer) loadData() {
	todoServiceServer.data = make(map[int32]*proto.TodoItem)
	// jsonStr := `{
	// 	"1": {
	// 		"id":"1",
	// 		"name":"Clean Bathroom",
	// 		"desc":"Clean the bathroom",
	// 		"days_left":"3"
	// 	},
	// 	"2": {
	// 		"id":"2",
	// 		"name":"Clean Kitchen",
	// 		"desc":"Clean the Kitchen",
	// 		"days_left":"14"
	// 	}
	// }`
	// json.Unmarshal([]byte(jsonStr), &todoServiceServer.data)
	// fmt.Printf("TestData %v", todoServiceServer.data)
	// var testData = map[string]proto.TodoItem{
	// 	"1": {
	// 		Id:1,
	// 		Name:"Clean Bathroom",
	// 		Desc:"Clean the bathroom",
	// 		DaysLeft:3,
	// 	},
	// 	"2": {
	// 		id:2,
	// 		name:"Clean Kitchen",
	// 		desc:"Clean the Kitchen",
	// 		daysleft:14,
	// 	},
	// }

}
