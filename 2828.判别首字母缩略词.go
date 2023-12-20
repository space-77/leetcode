/*
 * @lc app=leetcode.cn id=2828 lang=golang
 *
 * [2828] 判别首字母缩略词
 */

// @lc code=start
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i][0] != s[i] {
			return false
		}
	}

	return true
}

// @lc code=end

