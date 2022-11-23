package main

import (
	"fmt"
	"github.com/kocubinski/cosmos-adr-54/naive/foo/v1"
	foo_v1 "github.com/kocubinski/cosmos-adr-54/naive/foo/v1/api"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	//RegisterRoute
	//pb.RegisterRouteGuideServer(grpcServer, newServer())
	foo_v1.RegisterMsgServer(grpcServer, foo.NewMsgServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
