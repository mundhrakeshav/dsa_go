package arraysgfg

import (
	"testing"
)

// TestStackArray for testing Stack with Array
func TestEquilibriumPoint(t *testing.T) {

	t.Run("TestEquilibriumPoint", func(t *testing.T) {
		arr := []int{1, 3, 5, 2, 2}
		res := Equilibrium_Point(arr, uint(len(arr)))
		if res != 2 {
			t.Error("Invalid")
		}
		arr = []int{1}
		res = Equilibrium_Point(arr, uint(len(arr)))
		if res != 0 {
			t.Error("Invalid")
		}
		arr = []int{1, 3, 2, 4}
		res = Equilibrium_Point(arr, uint(len(arr)))
		if res != 2 {
			t.Error("Invalid")
		}
		
		arr = []int{1, 3, 2, 5}
		res = Equilibrium_Point(arr, uint(len(arr)))
		if res != -1 {
			t.Error("Invalid")
		}
		
	})

}
