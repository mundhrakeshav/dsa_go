package hash_test

import (
	hash "mundhrakeshav/dsa/gfg/hash"
	"reflect"
	"testing"
)

func TestDistinctElementsInAWindow(t *testing.T) {

	t.Run("TestDistinctElementsInAWindow", func(t *testing.T) {
		arr1 := []int{1, 2, 1, 3, 4, 2, 3}
		reflect.DeepEqual(hash.DistinctElementsInAWindow(arr1, 4), []int{3, 4, 4, 3})
	})
	
	t.Run("TestDistinctElementsInAWindow", func(t *testing.T) {
		arr1 := []int{1, 2, 4, 4}
		reflect.DeepEqual(hash.DistinctElementsInAWindow(arr1, 2), []int{2, 2, 1})
	})
}
