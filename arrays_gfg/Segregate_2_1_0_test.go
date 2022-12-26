package arraysgfg

import (
	"testing"
)

// TestStackArray for testing Stack with Array
func TestLeadersInAnArray(t *testing.T) {

	t.Run("TestLeadersInAnArray", func(t *testing.T) {
		arr := []int{16, 17, 4, 3, 5, 2};
		res := Leaders_In_A_Array(arr)

		if res[0] != 2 && res[1] != 5 && res[3] != 17 {
			t.Error("Invalid")
		}

	})

}
