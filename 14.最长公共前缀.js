/*
 * @lc app=leetcode.cn id=14 lang=javascript
 *
 * [14] 最长公共前缀
 */

// @lc code=start
/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  let res = ''
  for (let i = 0; i < strs[0].length; i++) {
    if (strs.length == strs.filter((e) => e[i] == strs[0][i]).length) res += strs[0][i]
    else return res
  }
  return res
}
// @lc code=end
