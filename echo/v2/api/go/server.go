package _go

import (
	"context"

	foov1 "github.com/kocubinski/cosmos-adr-54/echo/v1/foo"
	fooapi "github.com/kocubinski/cosmos-adr-54/echo/v2/api/gen"
	foov2 "github.com/kocubinski/cosmos-adr-54/echo/v2/foo"
)

var (
	_ fooapi.MsgServer = fooMsgServer{}
)

type fooMsgServer struct {
	keeperV2 foov2.Keeper
	keeperV1 foov1.Keeper
}

func (m fooMsgServer) DoSomething(_ context.Context, something *fooapi.MsgDoSomething) (*fooapi.MsgDoSomethingResponse, error) {
	// here we must run some validation to find out if it is v1 or v2 by introspecting the request
	return &fooapi.MsgDoSomethingResponse{Success: true},
		m.keeperV2.DoSomething(something.Sender, something.Amount, something.Condition)
}

func NewFooMsgServer() fooapi.MsgServer {
	return fooMsgServer{keeperV2: foov2.NewKeeper(), keeperV1: foov1.NewKeeper()}
}
