package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
)

type SpanData struct {
	price, span uint
}

func StockSpan(prices []uint) (spans []uint) {
	stk := stack.CreateNew[SpanData]()
	for _, v := range prices {
		if stk.IsEmpty() {
			stk = stk.Push(SpanData{
				price: v,
				span:  1,
			})
			spans = append(spans, 1)
		} else {
			total := uint(1)
			for !stk.IsEmpty() && stk.StackPeek().price < v {
				var pop SpanData
				pop, stk = stk.Pop()
				total = total + pop.span
			}
			stk = stk.Push(SpanData{
				price: v,
				span:  total,
			})
			spans = append(spans, total)
		}
	}
	return spans
}
