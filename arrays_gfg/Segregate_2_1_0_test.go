package arraysgfg

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestSegregate_2_1_0(t *testing.T) {

	t.Run("TestSegregate_2_1_0", func(t *testing.T) {
		arr := []int{2, 1, 0, 2, 1, 1, 1, 0, 1, 2, 2, 1, 2, 0};
		res := Segregate_2_1_0(arr)
		fmt.Println(res)

	})

}