package main

import (
	"context"
	"flag"
	"github.com/harry/grpcequtation/pkg/api"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	a, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	c, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		log.Fatal(err)
	}
	connection, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewSolverClient(connection)
	res, err := client.Solve(context.Background(), &api.SolveRequest{A: int32(a), B: int32(b), C: int32(c)})
	log.Println(res.GetAnswer())
}
