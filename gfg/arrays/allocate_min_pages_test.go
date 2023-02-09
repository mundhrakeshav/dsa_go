package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

func TestAllocateMinPages(t *testing.T) {

	t.Run("TestAllocateMinPages", func(t *testing.T) {
		arr := []int{12, 37, 64, 90}
		res := arraysgfg.AllocateMinPages(arr, 4, 2,)
		if res != 113 {
			t.Error("Invalid")
		}
	})
	
	t.Run("TestAllocateMinPages", func(t *testing.T) {
		arr := []int{15,17,20}
		res := arraysgfg.AllocateMinPages(arr, 3, 2,)
		if res != 32 {
			t.Error("Invalid")
		}
	})
}
