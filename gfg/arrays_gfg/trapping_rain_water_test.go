package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays_gfg"
	"testing"
)

func TestTrappingRainWater(t *testing.T) {

	// t.Run("TestTrappingRainWater", func(t *testing.T) {
	// 	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1 };
	// 	i := arraysgfg.TrappingRainWaterBruteForce(arr, len(arr))
	// 	if i!= 6{
	// 		t.Error("Invalid")
	// 	}
	// })
	
	// t.Run("TestTrappingRainWater", func(t *testing.T) {
	// 	arr := []int{3, 0, 2, 0, 4};
		
	// 	res := arraysgfg.TrappingRainWaterPreProcessing(arr, len(arr))
	// 	if res!= 7{
	// 		t.Error("Invalid")
	// 	}
	// })
	// t.Run("TrappingRainWaterStackApproach", func(t *testing.T) {
	// 	arr := []int{3, 0, 2, 0, 4};
		
	// 	res := arraysgfg.TrappingRainWaterStackApproach(arr, len(arr))
	// 	if res!= 7{
	// 		t.Error("Invalid")
	// 	}
	// 	arr = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1 };
	// 	res = arraysgfg.TrappingRainWaterStackApproach(arr, len(arr))
	// 	if res!= 6{
	// 		t.Error("Invalid")
	// 	}
		
	// })
	t.Run("TestTrappingRainWater", func(t *testing.T) {
		arr := []int{3, 0, 2, 0, 4};
		
		res := arraysgfg.TrappingRainWater2PointerApproach(arr, len(arr))
		
		if res!= 7{
			t.Error("Invalid")
		}
		arr = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1 };
		res = arraysgfg.TrappingRainWater2PointerApproach(arr, len(arr))

		if res!= 6{
			t.Error("Invalid")
		}
		
	})

}
