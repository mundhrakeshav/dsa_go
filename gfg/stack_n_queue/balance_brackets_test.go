package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"testing"
)

func TestBalancedBrackets(t *testing.T) {
	tests := []struct {
        expr     string
        expected bool
    }{
        {"()", true},
        {"(())", true},
        {"()()", true},
        {"(()())", true},
        {"(()))", false},
        {"((())", false},
        {"{[()]}", true},
        {"{[(])}", false},
        {"{{[[(())]]}}", true},
        {"{{[[(())]]}}}", false},
        {"[()]{}{[()()]()}", true},
        {"[(])", false},
    }

	for _, test := range tests {
        result := stack_n_queue.BalancedBrackets(test.expr)
        if result != test.expected {
            t.Errorf("BalancedBrackets(%q) = %v, expected %v", test.expr, result, test.expected)
        }
    }

}
