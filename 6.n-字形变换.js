/*
 * @lc app=leetcode.cn id=6 lang=javascript
 *
 * [6] N 字形变换
 */

// @lc code=start
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function (s, numRows) {
  if (s.length <= 1 || numRows == 1) return s

  const arr = new Array(numRows).fill('')
  let down = false
  let index = 0
  const len = Math.min(s.length, numRows)

  for (let i = 0; i < len; i++) {
    arr[index] += s[i]
    if (index == 0 || index == numRows - 1) down = !down
    index += down ? 1 : -1
  }

  return arr.join('')
}
// @lc code=end

const res = convert('PAYPALISHIRING', 4)
console.log(res === 'PINALSIGYAHRPI')
