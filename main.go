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

	fmt.Println(str)

	var opts []grpc.ServerOption
	// Checkthat we can listen on a port. Panic if there is an error
	fmt.Printf("Listening on %s\n", os.Getenv("GRPC_ADDR"))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		panic(err)
	}
	fmt.Printf("listener info: %v\n", lis.Addr())
	// Create a new gprcServer
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterTodoServiceServer(grpcServer, service.NewTodoItemServiceImpl())
	fmt.Printf("ServiceInfo: %v\n", grpcServer.GetServiceInfo())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
