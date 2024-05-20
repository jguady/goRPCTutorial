package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/jguady/goRPCTutorial/proto"
	"github.com/jguady/goRPCTutorial/service"
	"gotest.tools/v3/assert"
)

func TestOutput(t *testing.T) {

	// if expected, actual := "Hello World", messageOutput("Hello World"); actual != expected {
	// 	t.Errorf("The main function provided an actual of %q but %q was expected.", actual, expected)
	// }
}

func TestCreateItem(t *testing.T) {
	var a = int32(32)

	newId := rand.Int31()
	var expected = &proto.TodoItem{
		Id:       newId,
		Name:     "Foo",
		Desc:     "bar",
		DaysLeft: &a,
	}

	createItem := proto.CreateTodoItemRequest{
		Name:     "Foo",
		Desc:     "bar",
		DaysLeft: &a,
	}

	service := service.NewTodoItemServiceImpl()
	actual, err := service.CreateTodoItem(context.Background(), &createItem)
	// actual = expected
	if err != nil {
		println(err)
	}

	assert.Assert(t, actual.Id != 0)
	assert.Assert(t, expected.Name == actual.Name)
	assert.Assert(t, expected.DaysLeft == actual.DaysLeft)
	assert.Assert(t, expected.Desc == actual.Desc)

	t.Logf("Expected is : %v", expected)
	t.Logf("Actual   is : %v", actual)
}

func TestGetItem(t *testing.T) {

	service := service.NewTodoItemServiceImpl()
	var a = int32(16)

	createItem := proto.CreateTodoItemRequest{
		Name:     "Foo",
		Desc:     "bar",
		DaysLeft: &a,
	}

	expected, _ := service.CreateTodoItem(context.Background(), &createItem)

	actual, _ := service.GetTodoItem(context.Background(), &proto.GetTodoItemRequest{Id: expected.Id})
	fmt.Printf("Expected %v\n", expected)
	fmt.Printf("Actual %v\n", actual)
	assert.Assert(t, expected.Id == actual.Id)

}
