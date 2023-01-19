package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

func TestStockBuySell(t *testing.T) {

	t.Run("TestStockBuySell", func(t *testing.T) {
		arr := []int{100, 180, 260, 310, 40, 535, 695}
		res := arraysgfg.Stock_buy_sell(arr)
		if res!= 865 {
			t.Error("Invalid")
		}
		arr = []int{1, 4, 20, 3, 10, 5};
		res = arraysgfg.Stock_buy_sell(arr)
		if res!= 26 {
			t.Error("Invalid")
		}
		arr = []int {4, 2, 2, 2, 4};
		res = arraysgfg.Stock_buy_sell(arr)
		if res!= 2 {
			t.Error("Invalid")
		}

	})

}
