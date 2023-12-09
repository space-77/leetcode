/*
 * @lc app=leetcode.cn id=4 lang=javascript
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  let len1 = nums1.length
  let len2 = nums2.length
  let maxLen = len1 + len2
  let left = Math.floor((maxLen + 1) / 2)
  let right = Math.floor((maxLen + 2) / 2)
  return (findkth(nums1, 0, len1 - 1, nums2, 0, len2 - 1, left) + findkth(nums1, 0, len1 - 1, nums2, 0, len2 - 1, right)) / 2
}

function findkth(arr1, start1, end1, arr2, start2, end2, k) {
  let n = end1 - start1 + 1
  let m = end2 - start2 + 1
  if (n > m) return findkth(arr2, start2, end2, arr1, start1, end1, k)
  if (n === 0) return arr2[start2 + k - 1]
  if (k === 1) return Math.min(arr1[start1], arr2[start2])
  let i = start1 + Math.min(n, Math.floor(k / 2)) - 1
  let j = start2 + Math.min(m, Math.floor(k / 2)) - 1
  if (arr1[i] > arr2[j]) {
    return findkth(arr1, start1, end1, arr2, j + 1, end2, k - (j - start2 + 1))
  } else {
    return findkth(arr1, i + 1, end1, arr2, start2, end2, k - (i - start1 + 1))
  }
}
// @lc code=end

console.log(findMedianSortedArrays([1, 4, 44, 91, 89], [2, 5, 1000]))
// [(1, 2, 3, 7)]
