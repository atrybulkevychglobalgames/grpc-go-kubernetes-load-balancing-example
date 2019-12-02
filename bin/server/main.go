package main

import (
	"fmt"
	"github.com/dvcdsys/newinfragrpc/app/responder"
	service2 "github.com/dvcdsys/newinfragrpc/app/service"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := service2.ResponderService{
		Node: uuid.New().String(),
	}

	grpcServer := grpc.NewServer()
	responder.RegisterResponderServiceServer(grpcServer, &service)

	// start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Println("Server started")
}
