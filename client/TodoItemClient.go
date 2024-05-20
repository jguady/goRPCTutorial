package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/jguady/goRPCTutorial/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:50051", "The server address")
)

func randomTask() string {
	tasks := []string{
		"Wash dishes",
		"Clean the kitchen",
		"Vacuum the house",
		"Do the laundry",
		"Take out the trash",
		"Mow the lawn",
		"Grocery shopping",
		"Clean the bathroom",
		"Organize the garage",
		"Walk the dog",
		"Water the plants",
		"Dust the furniture",
		"Change the bed sheets",
		"Sweep the floors",
		"Cook dinner",
		"Wash the car",
		"Feed the pets",
		"Pay the bills",
		"Clean the windows",
		"Iron the clothes",
	}

	// Generate a random index
	randomIndex := rand.Intn(len(tasks))

	// Select a random task from the list
	randomTaskName := tasks[randomIndex]
	return randomTaskName
}

func main() {
	flag.Parse()
	//DialOption configures how we set up the connection.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Printf("Trying to connect to: %s\n", *serverAddr)
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(conn.CanonicalTarget())
	defer conn.Close()
	fmt.Println("Hello Client World")
	client := proto.NewTodoServiceClient(conn)

	createTodoItem(client)
	createTodoItem(client)
	createTodoItem(client)

	listTodoItems(client)
	item, _ := getTodoItem(client)
	updateTodoItem(client, item)

	listTodoItems(client)

	deleteTodoItem(client)
	listTodoItems(client)

}

func getTodoItem(client proto.TodoServiceClient) (*proto.TodoItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	item, err := client.GetTodoItem(ctx, &proto.GetTodoItemRequest{Id: 1})
	if err != nil {
		log.Fatalf("I've died getting item %d", 1)
		return nil, err
	}
	fmt.Printf("I GOT a Todo Item %v\n", item)
	return item, nil
}

func deleteTodoItem(client proto.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	itemToDelete, _ := client.GetTodoItem(ctx, &proto.GetTodoItemRequest{Id: 3})
	client.DeleteTodoItem(ctx, &proto.DeleteTodoItemRequest{Item: itemToDelete})

}

func updateTodoItem(client proto.TodoServiceClient, updateItem *proto.TodoItem) {

	//Setup Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Modify Item
	updatedItem := *updateItem
	updatedItem.Name = fmt.Sprintf("I have you now TK-%d", rand.Int31())
	newItem, err := client.UpdateTodoItem(ctx, &proto.UpdateTodoItemRequest{Id: updatedItem.Id, Item: &updatedItem})
	if err != nil {
		log.Fatalf("I've died Updating an item %d", updateItem.Id)
	}
	fmt.Printf("I Updated this item %v\n to this item %v\n", updateItem.Name, newItem.Name)

	// item, err := client.GetTodoItem(ctx, &proto.GetTodoItemRequest{Id: 1})
	// if err != nil {
	// 	log.Fatalf("I've died getting item %d", 1)
	// 	return nil, err
	// }
	// fmt.Printf("I GOT a Todo Item %v\n", item)
	// return item, nil
}

func listTodoItems(client proto.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	strim, err := client.ListTodoItems(ctx, &proto.ListTodoItemFilterRequest{})
	if err != nil {
		log.Fatalln("I've died listing items")
	}
	for {
		todoItem, err := strim.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("I died listing items two")
		}
		fmt.Printf("I'm a Todo Item %v\n", todoItem)
	}
}

func createTodoItem(client proto.TodoServiceClient) {

	createRequest := proto.CreateTodoItemRequest{Name: randomTask()}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item, err := client.CreateTodoItem(ctx, &createRequest)
	if err != nil {
		log.Fatalf("I died creating items, %v", err)
	}
	fmt.Printf("Item to Create %v\n", &item)
}
