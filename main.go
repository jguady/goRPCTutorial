package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jguady/goRPCTutorial/proto"
	"github.com/jguady/goRPCTutorial/service"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	str := "Hello World, setting context on signals SIGINT and SIGTERM"
	fmt.Println(str)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//Setup options for server.
	var opts []grpc.ServerOption
	// Checkthat we can listen on a port. Panic if there is an error
	fmt.Printf("Listening on %s\n", os.Getenv("GRPC_ADDR"))

	//Setup errorGroup
	errWg, errCtx := errgroup.WithContext(ctx)

	//Create a go Routine for the error(wait)Group, inside it creating an anonymous function
	grpcServer := grpc.NewServer(opts...)
	errWg.Go(func() error {
		// start the server
		listenerConfig := net.ListenConfig{}
		lis, err := listenerConfig.Listen(errCtx, "tcp", fmt.Sprintf("localhost:%d", 50051))
		if err != nil {
			return err
		}

		proto.RegisterTodoServiceServer(grpcServer, service.NewTodoItemServiceImpl())
		fmt.Printf("ServiceInfo: %v\n", grpcServer.GetServiceInfo())
		err = grpcServer.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})

	errWg.Go(func() error {
		<-errCtx.Done()
		grpcServer.GracefulStop()
		return nil
	})
	//wait for them to complete
	err := errWg.Wait()

	if err == context.Canceled || err == nil {
		fmt.Println("\nGracefully Stopping Server.\n")
	} else {
		fmt.Println(err)
	}
	// listenerConfig := net.ListenConfig{}
	// lis, err := listenerConfig.Listen()
	// // net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("listener info: %v\n", lis.Addr())
	// // Create a new gprcServer

}
