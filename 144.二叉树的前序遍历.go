package main

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) (arr []int) {
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node.Val)
		walk(node.Left)
		walk(node.Right)

	}
	walk(root)
	return
}

// @lc code=end
