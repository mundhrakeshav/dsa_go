package strings_test

import (
	strings "mundhrakeshav/dsa/gfg/strings"
	"testing"
)

func TestFindSubStr(t *testing.T) {

	t.Run("TestFindSubStr", func(t *testing.T) {
		s1 := "et"
		s2 := "omnisetomnis."
		x := strings.Check_Substr(s1, s2)
		if x != 5 {
			t.Error()
		}
	})
	
	// t.Run("TestFindSubStr", func(t *testing.T) {
	// 	str := "ABDEFGABEF."
	// 	x := strings.Longest_Unique_Substring(str)
	// 	if x != 6 {
	// 		t.Error()
	// 	}
	// })
	
	// t.Run("TestFindSubStr", func(t *testing.T) {
	// 	str := "BBBB."
	// 	x := strings.Longest_Unique_Substring(str)
	// 	if x != 1 {
	// 		t.Error()
	// 	}
	// })
	
	// t.Run("TestFindSubStr", func(t *testing.T) {
	// 	str := "GEEKSFORGEEKS."
	// 	x := strings.Longest_Unique_Substring(str)
	// 	if x != 7 {
	// 		t.Error()
	// 	}
	// })

}
