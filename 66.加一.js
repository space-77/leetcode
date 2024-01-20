/*
 * @lc app=leetcode.cn id=66 lang=javascript
 *
 * [66] 加一
 */

// @lc code=start
/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
  const len = digits.length

  for (let i = len - 1; i >= 0; i--) {
    digits[i] += 1
    digits[i] %= 10
    if (digits[i] !== 0) return digits
  }

  const res = Array(len+1).fill(0)
  res[0] = 1
  return res
};
// @lc code=end

