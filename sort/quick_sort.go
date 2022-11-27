package sort
func partition(_arr []int, _low, _high int) int {
    _pivotElement := _arr[_high];
    _pi := _low - 1;

    for _pj := _low; _pj < _high; _pj++ {
        if _arr[_pj] < _pivotElement {
            _pi++;
            _arr[_pi], _arr[_pj] = _arr[_pj], _arr[_pi]
        }
    }

    _arr[_pi + 1], _arr[_high] = _arr[_high], _arr[_pi + 1]
    return _pi + 1;
}

func QuickSortRange(_arr []int, _low, _high int)  {
    if _low >= _high {
        return
    }
    _pivot := partition(_arr, _low, _high);
    QuickSortRange(_arr, _low, _pivot - 1)
    QuickSortRange(_arr, _pivot + 1, _high)
}

func QuickSort(arr []int) []int {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}