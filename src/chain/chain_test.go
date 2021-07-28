package chain

import (
	"context"
	"testing"

	"github.com/kasiss/fncomposer/src/caller"
)

func TestChainRun(t *testing.T) {
	ch := NewFnChain(5)
	c5 := caller.NewCaller(func(c context.Context) { println(5); <-c.Done() }, 5)
	ch.Append(c5)
	c4 := caller.NewCaller(func(c context.Context) { println(4); <-c.Done() }, 4)
	ch.Append(c4)
	c3 := caller.NewCaller(func(c context.Context) { println(3); <-c.Done() }, 3)
	ch.Append(c3)
	c2 := caller.NewCaller(func(c context.Context) { println(2); <-c.Done() }, 2)
	ch.Append(c2)
	c1 := caller.NewCaller(func(c context.Context) { println(1); <-c.Done() }, 1)
	ch.Append(c1)

	ch.Run()
}
