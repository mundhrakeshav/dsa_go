/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package leetcode

// @lc code=start
func removeDuplicates(nums []int) int {
	i := 1

	for j := i; j < len(nums); j++ {
		if nums[j] > nums[i-1] {
			nums[i] = nums[j]
			i++;
		}
	}

	return i
}

// @lc code=end
