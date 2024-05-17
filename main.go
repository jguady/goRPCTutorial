package main

import (
	"fmt"
	"net"
	"os"

	"github.com/jguady/goRPCTutorial/proto"
	"github.com/jguady/goRPCTutorial/service"

	"google.golang.org/grpc"
)

func main() {
	str := "Hello World"

	fmt.Println(messageOutput(str))

	var opts []grpc.ServerOption
	// Checkthat we can listen on a port. Panic if there is an error
	lis, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}
	// Create a new gprcServer
	grpcServer := grpc.NewServer(opts...)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	proto.RegisterTodoServiceServer(grpcServer, service.NewTodoItemServiceImpl())

}

func messageOutput(str string) string {
	return str
}
