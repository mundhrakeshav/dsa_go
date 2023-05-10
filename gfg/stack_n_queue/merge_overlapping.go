package stacknqueue

import (
	"math"
	stack "mundhrakeshav/dsa/DS/stack"
	"sort"
)

type Interval struct {
	Start, End uint
}

func MergeOverlapping(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	stk := stack.CreateNew[Interval]()
	for i := len(intervals) - 1; i >= 0; i-- {
		interval := intervals[i]
		if stk.IsEmpty() {
			stk = stk.Push(interval)
		} else if top := stk.StackPeek(); interval.End >= top.Start {
			var pop Interval
			pop, stk = stk.Pop()
			pop = Interval{
				Start: uint(math.Min(float64(interval.Start), float64(pop.Start))),
				End:   uint(math.Max(float64(interval.End), float64(pop.End))),
			}
			stk = stk.Push(pop)
		} else {
			stk = stk.Push(interval)
		}
	}
	return stk
}
