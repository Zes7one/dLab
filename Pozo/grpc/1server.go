// Package main implements a server for Greeter service.
package main

import (
	// "context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Zes7one/dLab/Pozo/grpc/consulta"

	// pb "google.golang.org/grpc/examples/helloworld/h"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := consulta.Server{}

	grpcServer := grpc.NewServer()

	consulta.RegisterConsultaServiceServer(grpcServer, &s)

	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
