/*
 * @lc app=leetcode.cn id=7 lang=javascript
 *
 * [7] 整数反转
 */

// @lc code=start
/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
  let num = 0
  while (x !== 0) {
    num = num * 10 + (x % 10)
    if(num > Math.pow(2, 31) - 1 || num < Math.pow(-2, 31)) return 0;
    x = ~~(x / 10)
  }

  return num
}

reverse()
// @lc code=end
