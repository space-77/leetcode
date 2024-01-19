package main

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) (arr []int) {
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}

		walk(node.Left)
		walk(node.Right)
		arr = append(arr, node.Val)
	}

	walk(root)

	return

}

// @lc code=end
