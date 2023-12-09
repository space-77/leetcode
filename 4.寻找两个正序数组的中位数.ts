/*
 * @lc app=leetcode.cn id=4 lang=typescript
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  const maxLen = nums1.length + nums2.length
  const arr: number[] = []

  // 合并 数组并排序
  for (let i = 0, j = 0, k = 0; i < maxLen; i++) {
    if (nums1[j] < nums2[k]) arr.push(nums1[j++])
    else arr.push(nums2[k++])
  }
}
// @lc code=end
