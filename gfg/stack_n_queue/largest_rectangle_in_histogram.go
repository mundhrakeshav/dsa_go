package stacknqueue

import "math"

func LargestRectangleInHistogram(arr []uint) uint {
	// {2, 1, 5, 6, 2, 3, 1}
	// {1, 7, 4, 4, 6, 6, 7} nse
	// {-1, -1, 1, 2, 1, 4, -1} pse
	nse := NSE(arr)
	maxArea := uint(0)
	for i, v := range nse {
		if v == -1 {
			nse[i] = len(arr);
		}
	}
	pse := PSE(arr)
	
	//-----------------------------/
	for i, v := range arr {
		width := nse[i] - pse[i] - 1
		area := v * uint(width)
		maxArea = uint(math.Max(float64(area), float64(maxArea)))
	}

	return maxArea;

}