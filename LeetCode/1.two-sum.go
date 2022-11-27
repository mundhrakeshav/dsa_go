/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetcode

import (
	"sort"
)

func TwoSumNaive1(nums []int, target int) []int {
    sort.Ints(nums); // If array is sorted this approach can be implemented
    i := 0;
    j := len(nums) - 1;
    
    for i < j {
        if nums[i] + nums[j] > target {
            j--;
        } else if(nums[i] + nums[j] < target) {
            i++;
        } else {
            return []int {i, j}
        }
    }
    return []int{};
}

type Record struct {
    isThere bool;
    index   int;
}
func TwoSum1(nums []int, target int) []int {
    var record map[int]Record = make(map[int]Record);

    for i := 0; i < len(nums); i++ {
        diff := target - nums[i];
        if record[diff].isThere {
            return []int{record[diff].index, i}
        }
        record[nums[i]] = Record{
            isThere: true,
            index: i,
        }
    }

    return []int{-1}
}

func TwoSum(nums []int, target int) []int {
    var record map[int]int = make(map[int]int);

    for i := 0; i < len(nums); i++ {
        diff := target - nums[i];
        x,exists := record[diff]
        if exists  {
            return []int{x, i}
        }
        record[nums[i]] = i;
    }

    return nil;
}