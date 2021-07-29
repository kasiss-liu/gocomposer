package chain

import (
	"context"
	"time"

	"github.com/kasiss-liu/gocomposer/src/caller"
)

type ChainContext struct {
	wait chan struct{}
	ctx  context.Context
}

func NewChainContext(ctx ...context.Context) *ChainContext {
	c := &ChainContext{}
	c.wait = make(chan struct{}, 1)
	c.wait <- struct{}{}
	if len(ctx) > 0 && ctx[0] != nil {
		c.ctx = ctx[0]
	} else {
		c.ctx = context.TODO()
	}
	return c
}

func (cc ChainContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}
func (cc ChainContext) Done() <-chan struct{} {
	return cc.wait
}
func (cc ChainContext) Value(key interface{}) interface{} {
	return cc.ctx.Value(key)
}
func (cc ChainContext) Err() error {
	return nil
}
func (cc ChainContext) Wait() {
	cc.wait <- struct{}{}
}

type Chain struct {
	Callers *caller.Callers
	Ctx     *ChainContext
}

func (ch *Chain) Append(c *caller.Caller) {
	ch.Callers.Add(c)
}

func (ch *Chain) Run() {
	ch.Callers.Sort()
	l := ch.Callers.Len()

	for i := 0; i < l; i++ {
		go ch.Callers.Value(i).Fn(ch.Ctx)
		ch.Ctx.Wait()
	}
}

func NewFnChain(cap ...int) *Chain {
	c := 5
	if len(cap) > 0 {
		c = cap[0]
	}
	ch := &Chain{}
	ch.Callers = caller.NewCallers(c)
	ch.Ctx = NewChainContext()
	return ch
}
