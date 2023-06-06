package arraysgfg

import (
	"math"
	stack "mundhrakeshav/dsa/ds/stack"
)

func TrappingRainWaterBruteForce(arr []int, n int) int {
	res := 0
	for i := 1; i < n-1; i++ {

		left_max := arr[i]
		for j := 0; j < i; j++ {
			if arr[j] > left_max {
				left_max = arr[j]
			}
		}
		right_max := arr[i]
		for j := i; j < n; j++ {
			if arr[j] > right_max {
				right_max = arr[j]
			}
		}

		if left_max < right_max {
			res += left_max - arr[i]
		} else {
			res += right_max - arr[i]
		}

	}
	return res
}

func TrappingRainWaterPreProcessing(arr []int, n int) int {
	left_max := make([]int, n)
	right_max := make([]int, n)
	res := 0
	max_left_yet := 0
	max_right_yet := 0
	for i := 0; i < n; i++ {
		if arr[i] > max_left_yet {
			max_left_yet = arr[i]
		}
		left_max[i] = max_left_yet
	}
	for i := n - 1; i >= 0; i-- {
		if arr[i] > max_right_yet {
			max_right_yet = arr[i]
		}
		right_max[i] = max_right_yet
	}

	for i := 0; i < n; i++ {
		res += int(math.Min(float64(right_max[i]), float64(left_max[i]))) - arr[i]
	}
	return res

}

func TrappingRainWaterStackApproach(arr []int, n int) int {
	_stack := stack.CreateNew[int]()
	res, _top := 0, 0
	for i := 0; i < n; i++ {
		//	if stack isn't empty	&& stack.top() > stack[current]
		for !_stack.IsEmpty() && arr[i] > arr[_stack.StackPeek()] {
			_top, _stack = _stack.Pop()
			if _stack.IsEmpty() {
				break
			}
			dist := i - _stack.StackPeek() - 1
			height := int(math.Min(float64(arr[i]), float64(arr[_stack.StackPeek()]))) - arr[_top]
			res += (height * dist)
		}
		_stack = _stack.Push(i)
	}

	return res
}

func TrappingRainWater2PointerApproach(arr []int, n int) int {
	l, r, left_max, right_max, res := 0, n -1, 0, 0, 0

	for l <= r {
		if arr[l] <= arr[r] {
			if arr[l] < left_max {
				res += left_max - arr[l]
			} else {
				left_max = arr[l]
			}
			l++
		} else {
			if arr[r] < right_max {
				res += right_max - arr[r]
			} else {
				right_max = arr[r]
			}
			r--
		}
	} 
	return res;
}
