package main

import (
	"fmt"
	"net"

	"github.com/kocubinski/cosmos-adr-54/echo/v1/api"
	fooapi "github.com/kocubinski/cosmos-adr-54/echo/v1/api/gen"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	fooapi.RegisterMsgServer(grpcServer, api.NewFooMsgServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
