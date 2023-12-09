/*
 * @lc app=leetcode.cn id=3 lang=javascript
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  if (s.length <=1 ) return s.length;
  let left = 0;
  let right = 1;
  let temp = "";
  let maxlen = 0;

  while (s.length > right) {
    temp = s.slice(left, right);

    if (temp.indexOf(s[right]) > -1) {
      left++;
    } else {
      right++;
      const max = right - left;
      maxlen = max > maxlen ? max : maxlen;
    }
  }

  return maxlen;
};

console.log(lengthOfLongestSubstring(" "));
// @lc code=end
