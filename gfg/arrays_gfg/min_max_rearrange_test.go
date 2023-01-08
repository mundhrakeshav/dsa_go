package arraysgfg_test

import (
	"fmt"
	arraysgfg "mundhrakeshav/dsa/gfg/arrays_gfg"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestMaxMinRearrange(t *testing.T) {

	t.Run("TestMaxMinFormArrayRearrange", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6}
		arr = arraysgfg.Min_Max_Rearrange(arr)
		fmt.Println(arr)
		// if res[0] != 2 && res[1] != 5 && res[3] != 17 {
		// 	t.Error("Invalid")
		// }

	})

}
