package stacknqueue_test

import (
	"fmt"
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"testing"
)

func TestGetMinStack(t *testing.T) {

	t.Run("Create New", func(t *testing.T) {
		getMinStk := stack_n_queue.CreateNewGetMinStack[int]()
		getMinStk = getMinStk.Push(12);
		fmt.Println(getMinStk.Min)
		getMinStk = getMinStk.Push(13);
		fmt.Println(getMinStk.Min)
		getMinStk = getMinStk.Push(11);
		fmt.Println(getMinStk.Min)
		getMinStk = getMinStk.Push(10);
		fmt.Println(getMinStk.Min)
		var pop int;
		pop, getMinStk = getMinStk.Pop();
		fmt.Println(getMinStk.Min, pop)
		pop, getMinStk = getMinStk.Pop();
		fmt.Println(getMinStk.Min, pop)
		pop, getMinStk = getMinStk.Pop();
		fmt.Println(getMinStk.Min, pop)
	})



}
