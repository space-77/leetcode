package main

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n <= 1 {
		return n
	}

	f1 := 0
	f2 := 1
	f3 := 1

	for i := 2; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}

	return f3
}

// @lc code=end
