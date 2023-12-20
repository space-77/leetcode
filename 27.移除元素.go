package main

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	length := len(nums)

	slow := 0
	first := 0

	for first < length {
		if nums[first] != val {
			nums[slow], nums[first] = nums[first], nums[slow]
			slow++
			first++
		} else {
			first++
		}
	}

	return slow

}

// @lc code=end
