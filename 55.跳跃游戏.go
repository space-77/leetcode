package main

import "fmt"

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	cover := 0

	for i := 0; i <= cover; i++ {
		if (i + nums[i]) > cover {
			cover = (i + nums[i])
		}

		if cover >= len(nums)-1 {
			return true
		}

	}

	return false
}

// @lc code=end

func main() {
	res := canJump([]int{0, 2, 3})
	fmt.Println(res)
}
