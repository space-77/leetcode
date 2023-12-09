/*
 * @lc app=leetcode.cn id=15 lang=javascript
 *
 * [15] 三数之和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  const list = []
  if (nums.length < 3) return []
  nums.sort((a, b) => a - b)

  // [-1,0,1,2,-1,-4]
  for (let i = 0; i < nums.length; i++) {
    let left = i + 1
    let right = nums.length - 1
    const num = nums[i]

    // 跳过 相同的
    if (num === nums[i - 1]) continue

    while (left < right) {
      const leftNum = nums[left]
      const rightNum = nums[right]
      const count = num + leftNum + rightNum

      if (count === 0) {
        list.push([num, leftNum, rightNum]) // 添加数据

        // 命中后，让 left 继续往右查找
        while (leftNum === nums[left + 1]) left++ // 跳过相同的数
        left++ // left 往右

        while (rightNum === nums[right - 1]) right-- // 跳过相同的数
        right-- // right 往左，因为命中了，右边已经没有合适的数了，只能往左
      } else if (count > 0) {
        // 数大了右边向左
        right--
      } else if (count < 0) {
        // 数小了，左边向右
        left++
      }
    }
  }
  return list
}
// @lc code=end
