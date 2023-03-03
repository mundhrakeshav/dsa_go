package strings_test

import (
	strings "mundhrakeshav/dsa/gfg/strings"
	"testing"
)

func TestUniqueSubstring(t *testing.T) {

	t.Run("TestUniqueSubstring", func(t *testing.T) {
		str := "omnisomniset."
		x := strings.Longest_Unique_Substring(str)
		if x != 5 {
			t.Error()
		}
	})
	
	t.Run("TestUniqueSubstring", func(t *testing.T) {
		str := "ABDEFGABEF."
		x := strings.Longest_Unique_Substring(str)
		if x != 6 {
			t.Error()
		}
	})
	
	t.Run("TestUniqueSubstring", func(t *testing.T) {
		str := "BBBB."
		x := strings.Longest_Unique_Substring(str)
		if x != 1 {
			t.Error()
		}
	})
	
	t.Run("TestUniqueSubstring", func(t *testing.T) {
		str := "GEEKSFORGEEKS."
		x := strings.Longest_Unique_Substring(str)
		if x != 7 {
			t.Error()
		}
	})

}
