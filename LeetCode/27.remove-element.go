/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
package leetcode

// @lc code=start
func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1

	for i < j {
		for nums[i] != val {
			i++
		}
		for nums[j] == val {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}

// @lc code=end
