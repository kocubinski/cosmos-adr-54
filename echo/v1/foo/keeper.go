package foo

type Keeper interface {
	DoSomething(sender string, amount uint64) error
}
