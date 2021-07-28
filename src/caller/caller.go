package caller

import (
	"context"
	"sort"
)

type Caller struct {
	Fn    func(context.Context)
	Score int
}

func NewCaller(fn func(context.Context), score int) *Caller {
	return &Caller{Fn: fn, Score: score}

}

type Callers []*Caller

func NewCallers(cap int) *Callers {
	cs := make(Callers, 0, cap)
	return &cs
}

func (cs *Callers) Add(c *Caller) {
	*cs = append(*cs, c)
}
func (cs Callers) Value(index int) *Caller {
	return cs[index]
}

func (cs Callers) Sort() {
	sort.Sort(cs)
}
func (cs Callers) Len() int {
	return len(cs)
}
func (cs Callers) Less(i, j int) bool {
	return cs[i].Score < cs[j].Score
}
func (cs Callers) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
