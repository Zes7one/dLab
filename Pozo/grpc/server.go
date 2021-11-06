package main

import (
	"log"
	"net"

	"github.com/Zes7one/dLab/Pozo/grpc/consulta"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Fallo al escuchar en puerto 9000: %v", err)
	}

	s := consulta.Server{}

	grpcServer := grpc.NewServer()

	consulta.RegisterConsultaServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir server grpc en puerto 9000: %v", err)
	}
}
