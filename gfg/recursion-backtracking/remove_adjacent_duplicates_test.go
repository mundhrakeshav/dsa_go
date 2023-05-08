package recursionbacktracking_test

import (
	"fmt"
	recursion_backtracking "mundhrakeshav/dsa/gfg/recursion-backtracking"
	"testing"
)

// Incomplete
func TestSmallestPrimeFactors(t *testing.T) {
	t.Run("RemoveAdjacentDuplicates", func(t *testing.T) {
		s := recursion_backtracking.RemoveAdjacentDuplicates("caaabbbaacdddd")
		fmt.Println(s)

	})
}
// c aaa bbb aa c dddd