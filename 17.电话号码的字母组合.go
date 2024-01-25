package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) (ans []string) {
	size := len(digits)
	if size == 0 {
		return
	}
	keyMap := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	textList := [][]byte{}
	for _, digit := range digits {
		texts, ok := keyMap[byte(digit)]
		if ok {
			textList = append(textList, texts)
		}
	}

	var backtrack func(list []byte, deep int)
	backtrack = func(list []byte, deep int) {
		if deep == size {
			ans = append(ans, string(list))
			return
		}

		nextDeep := deep + 1
		for _, v := range textList[deep] {
			backtrack(append(list, v), nextDeep)
		}
	}

	backtrack([]byte{}, 0)
	return
}

// @lc code=end
