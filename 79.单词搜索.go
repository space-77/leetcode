package main

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	colLen := len(board)
	rowLen := len(board[0])

	var find func(col int, row int, cur int) bool
	find = func(col int, row int, cur int) bool {
		if col < 0 || col >= colLen || row >= rowLen || row < 0 {
			return false
		}
		text := board[col][row]

		if text != word[cur] {
			return false
		}
		if cur == len(word)-1 {
			return true
		}

		var ep byte
		board[col][row] = ep

		nextCur := cur + 1
		res := find(col+1, row, nextCur) || find(col-1, row, nextCur) || find(col, row+1, nextCur) || find(col, row-1, nextCur)

		board[col][row] = text

		return res

	}

	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			res := find(i, j, 0)
			if res {
				return res
			}
		}
	}

	return false
}

// @lc code=end
