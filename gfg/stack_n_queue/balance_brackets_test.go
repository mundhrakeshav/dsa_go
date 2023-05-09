package stacknqueue_test

import (
	"fmt"
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	t.Run("TestBinarySearch", func(t *testing.T) {
		expr := "[()]{}{[()()]()}"
		isBalanced := stack_n_queue.BalancedBrackets(expr)
		fmt.Println(isBalanced)
	})

}
