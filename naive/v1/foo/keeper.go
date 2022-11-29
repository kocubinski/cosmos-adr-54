package foo

import (
	"context"
	"fmt"

	foo_v1 "github.com/kocubinski/cosmos-adr-54/naive/foo/api"
)

type Keeper interface {
	DoSomething(something foo_v1.MsgDoSomething) error
}

var (
	_ foo_v1.MsgServer = msgServer{}
	_ Keeper           = keeper{}
)

type msgServer struct {
	Keeper
}

type keeper struct{}

func (k keeper) DoSomething(something foo_v1.MsgDoSomething) error {
	fmt.Printf("DoSomething: %v\n", something)
	return nil
}

func (m msgServer) DoSomething(
	_ context.Context, something *foo_v1.MsgDoSomething,
) (*foo_v1.MsgDoSomethingResponse, error) {
	err := m.Keeper.DoSomething(*something)
	if err != nil {
		return &foo_v1.MsgDoSomethingResponse{Success: false}, err
	}
	return &foo_v1.MsgDoSomethingResponse{Success: true}, nil
}

func NewMsgServer() foo_v1.MsgServer {
	return msgServer{Keeper: keeper{}}
}
