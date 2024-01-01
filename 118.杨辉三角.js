/*
 * @lc app=leetcode.cn id=118 lang=javascript
 *
 * [118] 杨辉三角
 */
// var generate = function(numRows) {
//   const dp = [];
//   for (let i = 0; i < numRows; i++) {
//       dp[i] = Array(i+1).fill(1);
//       for (let j = 1; j < i; j++) {
//           dp[i][j]=dp[i-1][j-1]+dp[i-1][j];
//       }
//   }
//   return dp;
// };

// @lc code=start
/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
  const ans = [];
  ans.push([1]);
  for (let i = 1; i < numRows; i++) {
    const temp = new Array(i + 1);
    temp[0] = 1;
    temp[temp.length - 1] = 1;

    
    const len = Math.ceil(temp.length / 2)
    for (let j = 1, k = temp.length - 2; j < len; j++) {
      temp[j] = ans[i - 1][j - 1] + ans[i - 1][j];
      temp[k] = temp[j]
      k--
    }
    
    ans.push(temp);
  }
  
  return ans;
};
// @lc code=end

