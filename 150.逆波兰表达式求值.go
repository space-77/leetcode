package main

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}

	var getNums func() (val1 int, val2 int)
	getNums = func() (val1 int, val2 int) {
		i := len(stack) - 1
		j := len(stack) - 2
		val1 = stack[i]
		val2 = stack[j]
		stack = stack[:j]
		return
	}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		val, err := strconv.Atoi(token)
		var num int

		if err == nil {
			num = val
		} else {
			val1, val2 := getNums()
			if token == "+" {
				num = val1 + val2
			} else if token == "-" {
				num = val2 - val1
			} else if token == "*" {
				num = val1 * val2
			} else if token == "/" {
				num = val2 / val1
			}
		}

		stack = append(stack, num)

	}

	return stack[0]
}

// @lc code=end
