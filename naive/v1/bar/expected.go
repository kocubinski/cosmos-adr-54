package bar

import foo_v1 "github.com/kocubinski/cosmos-adr-54/naive/v1/foo/api"

type FooKeeper interface {
	DoSomething(something foo_v1.MsgDoSomething) error
}
