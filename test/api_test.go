package main

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/kocubinski/cosmos-adr-54/echo/v1/api"
	pbv1 "github.com/kocubinski/cosmos-adr-54/echo/v1/api/gen"
	apiv2 "github.com/kocubinski/cosmos-adr-54/echo/v2/api"
	pbv2 "github.com/kocubinski/cosmos-adr-54/echo/v2/api/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestFoo(t *testing.T) {
	v1Listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	require.NoError(t, err)
	v2Listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	require.NoError(t, err)

	v1Server := grpc.NewServer()
	v2Server := grpc.NewServer()
	defer v1Server.Stop()
	defer v2Server.Stop()
	pbv1.RegisterMsgServer(v1Server, api.NewFooMsgServer())
	pbv2.RegisterMsgServer(v2Server, apiv2.NewFooMsgServer())
	go func() {
		err = v1Server.Serve(v1Listener)
		if err != nil {
			t.Errorf("v1Server.Serve() error = %v", err)
		}
	}()
	go func() {
		err = v2Server.Serve(v2Listener)
		if err != nil {
			t.Errorf("v2Server.Serve() error = %v", err)
		}

	}()

	ctx := context.Background()
	v1Conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	v2Conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	require.NoError(t, err)

	v1Client := pbv1.NewMsgClient(v1Conn)
	v2Client := pbv2.NewMsgClient(v2Conn)

	// v1Client with v1 Server
	_, err = v1Client.DoSomething(ctx, &pbv1.MsgDoSomething{Sender: "v1", Amount: 1})
	require.NoError(t, err)

	// v1Client with v2 Server
	v2Out := new(pbv2.MsgDoSomethingResponse)
	err = v2Conn.Invoke(ctx, "/foo.v2.Msg/DoSomething", &pbv1.MsgDoSomething{Sender: "v1", Amount: 1}, v2Out)
	require.NoError(t, err)

	// v2Client with v1 Server
	v1Out := new(pbv1.MsgDoSomethingResponse)
	err = v1Conn.Invoke(ctx, "/foo.v1.Msg/DoSomething", &pbv2.MsgDoSomething{Sender: "v2", Amount: 2, Condition: "foo"}, v1Out)

	// v2Client with v2 Server
	_, err = v2Client.DoSomething(ctx, &pbv2.MsgDoSomething{Sender: "v1", Amount: 1, Condition: "foo"})
	require.NoError(t, err)
}
