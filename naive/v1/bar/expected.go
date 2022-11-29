package v1

import foo_v1 "github.com/kocubinski/cosmos-adr-54/naive/foo/v1/api"

type FooKeeper interface {
	DoSomething(something foo_v1.MsgDoSomething) error
}
