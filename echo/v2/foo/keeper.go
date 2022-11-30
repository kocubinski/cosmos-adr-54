package foo

import "fmt"

type Keeper interface {
	DoSomething(sender string, amount uint64) error
}

type keeper struct{}

func (k keeper) DoSomething(sender string, amount uint64) error {
	fmt.Printf("DoSomething: %v from %v\n", sender, amount)
	return nil
}

var _ Keeper = keeper{}

func NewKeeper() Keeper {
	return keeper{}
}
