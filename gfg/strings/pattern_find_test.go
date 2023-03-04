package strings_test

import (
	strings "mundhrakeshav/dsa/gfg/strings"
	"reflect"
	"testing"
)

func TestPattern_find(t *testing.T) {

	t.Run("TestPattern_find", func(t *testing.T) {
		txt := "omniset.omnis"
		pat := "omnis"
		x := strings.Pattern_find(txt, pat)
		if !reflect.DeepEqual(x, []int{0, 8}) {
			t.Error()
		}
	})
	
	t.Run("TestPattern_find", func(t *testing.T) {
		txt := "AABAACAADAABAABA"
		pat := "AABA"
		x := strings.Pattern_find(txt, pat)
		if !reflect.DeepEqual(x, []int{0, 9, 12}) {
			t.Error()
		}
	})
	
}
