package api

import (
	"context"

	fooapi "github.com/kocubinski/cosmos-adr-54/echo/api/go/gen"
	"github.com/kocubinski/cosmos-adr-54/echo/v1/foo"
)

var (
	_ fooapi.MsgServer = fooMsgServer{}
)

type fooMsgServer struct {
	foo.Keeper
}

func (m fooMsgServer) DoSomething(_ context.Context, something *fooapi.MsgDoSomething) (*fooapi.MsgDoSomethingResponse, error) {
	return &fooapi.MsgDoSomethingResponse{Success: true}, m.Keeper.DoSomething(something.Sender, something.Amount)
}

func NewFooMsgServer() fooapi.MsgServer {
	return fooMsgServer{Keeper: foo.NewKeeper()}
}
