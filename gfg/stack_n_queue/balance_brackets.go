package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
)

func BalancedBrackets(expr string) bool {
	stk := stack.CreateNew[rune]()
	for _, v := range expr {
		if isOpenBracket(v) {
			stk = stk.Push(v)
		} else {
			var pop rune;
			pop, stk = stk.Pop()
			if !isValidBracketPair(v, pop) {
				return false
			}
		}
	}
	return stk.IsEmpty();
	}

func isValidBracketPair(_symbol1, _symbol2 rune) bool {
	switch _symbol1 {
	case '}': if _symbol2 == '{' {return true} else {return false}
	case ']': if _symbol2 == '[' {return true} else {return false}
	case ')': if _symbol2 == '(' {return true} else {return false}
	default: return false
	}
}
func isOpenBracket(_symbol rune) bool {
	switch _symbol {
	case '{': return true
	case '[': return true
	case '(': return true
	default: return false
	}
}