package foo

import "fmt"

type Keeper interface {
	DoSomething(sender string, amount uint64, condition string) error
}

type keeper struct{}

func (k keeper) DoSomething(sender string, amount uint64, condition string) error {
	fmt.Printf("DoSomething: %v from %v with %v\n", sender, amount, condition)
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	return nil
}

var _ Keeper = keeper{}

func NewKeeper() Keeper {
	return keeper{}
}
