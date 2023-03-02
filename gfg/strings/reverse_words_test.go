package strings_test

import (
	strings "mundhrakeshav/dsa/gfg/strings"
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {

	t.Run("TestReverseWords", func(t *testing.T) {
		str := "Omnis omnis et."
		x := strings.Reverse_Words(str)
		if reflect.DeepEqual(x, []string{"et", "omnis", "Omnis"}) {
			t.Error()
		}
		
	})

}
