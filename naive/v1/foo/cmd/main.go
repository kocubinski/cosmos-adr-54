package main

import (
	"fmt"
	"net"

	"github.com/kocubinski/cosmos-adr-54/naive/foo"
	foo_v1 "github.com/kocubinski/cosmos-adr-54/naive/foo/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	foo_v1.RegisterMsgServer(grpcServer, foo.NewMsgServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
