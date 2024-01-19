package main

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) (arr []int) {
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		arr = append(arr, node.Val)
		walk(node.Right)
	}
	walk(root)
	return
}

// @lc code=end
