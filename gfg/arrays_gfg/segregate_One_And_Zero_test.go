package arraysgfg

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestSegregate0And1(t *testing.T) {

	t.Run("TestSegregate0And1", func(t *testing.T) {
		arr := []int{0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0}
		arr = Segregate_One_and_Zero(arr)
		fmt.Println(arr)
		

	})

}
