package main

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	length := len(nums)
	arr := make([]int, length)
	left := 0
	right := length - 1

	for i := right; i >= 0; i-- {
		if nums[right]+nums[left] > 0 {
			arr[i] = nums[right] * nums[right]
			right--
		} else {
			arr[i] = nums[left] * nums[left]
			left++
		}
	}

	return arr

}

// @lc code=end
