/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	slow, first, length := 0, 0, len(nums)
	for first < length {
		if nums[first] != 0 {
			nums[slow], nums[first] = nums[first], nums[slow]
			slow++
		}

		first++
	}
}

// @lc code=end

// 0,1,0,3,12
// 1,3,12,0,0