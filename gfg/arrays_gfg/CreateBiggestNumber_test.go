package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays_gfg"
	"testing"
)

func TestCreateBiggestNumber(t *testing.T) {

	t.Run("TestCreateBiggestNumber", func(t *testing.T) {
		arr := []string{"1", "34", "3", "98", "9", "76", "45", "4"}
		res := arraysgfg.Create_Biggest_Number(arr)
		if res != "998764543431" {
			t.Error("Invalid")
		}
		
	})

}
