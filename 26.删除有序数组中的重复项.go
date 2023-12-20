package main

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow := 0
	first := 1
	length := len(nums)

	if length <= 1 {
		return 1
	}

	for first < length {
		if nums[first] > nums[slow] {
			slow++
			nums[slow] = nums[first]
		}
		first++
	}

	return slow + 1
}

// @lc code=end
