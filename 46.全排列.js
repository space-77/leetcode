/*
 * @lc app=leetcode.cn id=46 lang=javascript
 *
 * [46] 全排列
 */

// @lc code=start

/**
 *
 * @param {Array} list
 * @param {Array} nums
 * @param {Array} arr
 */
function backtrack(list, nums, arr) {
  if (nums.length === arr.length) {
    return list.push([...arr]);
  }

  for (const num of nums) {
    if (arr.includes(num)) continue;

    arr.push(num);
    backtrack(list, nums, arr);
    arr.pop();
  }
}

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function (nums) {
  const list = [];

  backtrack(list, nums, []);

  return list;
};
// @lc code=end
