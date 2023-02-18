package hash_test

import (
	hash "mundhrakeshav/dsa/gfg/hash"
	"testing"
)

func TestIfArrayCanBeDividedIntoPairsDivByK(t *testing.T) {

	t.Run("TestIfArrayCanBeDividedIntoPairsDivByK", func(t *testing.T) {
		arr := []int{9, 7, 5, 3}
		res := hash.CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr, 3)
		if !res { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestIfArrayCanBeDividedIntoPairsDivByK", func(t *testing.T) {
		arr := []int{9, 7, 5, 3}
		res := hash.CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr, 6)
		if !res { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestIfArrayCanBeDividedIntoPairsDivByK", func(t *testing.T) {
		arr := []int{92, 75, 65, 48, 45, 35}
		res := hash.CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr, 10)
		if !res { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestIfArrayCanBeDividedIntoPairsDivByK", func(t *testing.T) {
		arr := []int{91, 74, 66, 48}
		res := hash.CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr, 10)
		if res { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestIfArrayCanBeDividedIntoPairsDivByK", func(t *testing.T) {
		arr := []int{-1, 1, -2, 2, -3, 3}
		res := hash.CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr, 10)
		if !res { 
			t.Error("Invalid")
		}
	})
	
}
