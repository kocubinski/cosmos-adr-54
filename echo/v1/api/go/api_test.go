package api_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/kocubinski/cosmos-adr-54/echo/v1/api"
	fooapi "github.com/kocubinski/cosmos-adr-54/echo/v1/api/gen"
	"google.golang.org/grpc"
)

func TestFoo(t *testing.T) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	fooapi.RegisterMsgServer(grpcServer, api.NewFooMsgServer())
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	//ctx := context.Background()
	//conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	//err := conn.Invoke(ctx, "/foo.Msg/DoSomething", in, out)

	grpcServer.Stop()
}
