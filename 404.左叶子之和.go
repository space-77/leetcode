package main

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		left := node.Left
		if left != nil && left.Left == nil && left.Right == nil {
			sum += left.Val
		}

		walk(node.Left)
		walk(node.Right)
	}

	walk(root)

	return sum
}

// @lc code=end
