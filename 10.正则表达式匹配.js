/*
 * @lc app=leetcode.cn id=10 lang=javascript
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */

// 'aab'
// /c*a*b/
var isMatch = function (s, p) {
  let i = 0
  let pi = 0
  while (i < s.length - 1 || pi < p.length - 1) {
    const str = s[i]
    const strp = p[pi]

    if (str === strp || strp === '.') {
      i++
      pi++
      continue
    } else if (strp === '*') {
      const prePS = p[pi - 1]

      if (prePS === '.') {
        i++
        continue
      }
      if (str === prePS) {
        i++
        continue
      } else {
        // i++
        pi++
        continue
      }
    } else {
      // // 尝试 移动 p
      // let pii = pi
      // let pflag = false
      // while (pii < p.length - 1 && !pflag) {
      //   if (p[pii] === str  || p[pii] === '.') {
      //     pflag = true
      //     continue
      //   }
      //   pii++
      // }
      // if (pflag) {
      //   pi = pii
      //   continue
      // }

      // // 尝试移动 i
      // let ii = i
      // let sflag = false
      // while (ii < s.length - 1 && !sflag) {
      //   if (s[ii] === strp || s[ii] === '.') {
      //     sflag = true
      //     continue
      //   }
      //   ii++
      // }

      // if (pflag) {
      //   i = ii
      //   continue
      // }

      // i++
      // return false
    }
  }

  return i === s.length - 1 && pi === p.length - 1
}
// @lc code=end

// console.log(isMatch('aab', 'c*a*b'))
console.log(isMatch('ippi', 'p*.'));
// console.log(isMatch('mississippi', 'mis*is*p*.'));
