package main

import (
	"github.com/harry/grpcequtation/pkg/api"
	"github.com/harry/grpcequtation/pkg/solver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	microService := &solver.GRPCServer{}
	api.RegisterSolverServer(server, microService)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
