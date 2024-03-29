/*
 * @lc app=leetcode.cn id=114 lang=javascript
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {void} Do not return anything, modify root in-place instead.
 */
var flatten = function (root) {
  const list = []

  function walk(node, nodeList) {
    if (!node) return
    list.push(node)
    walk(node.left, nodeList)
    walk(node.right, nodeList)
  }
  walk(root, list)

  for (let i = 0; i < list.length; i++) {
    const node = list[i]
    const nextNode = list[i + 1] || null
    node.left = null
    node.right = nextNode
  }
}
// @lc code=end
