package main

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	var travese func(node *TreeNode) bool
	prev := math.MinInt

	travese = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		left := travese(node.Left)

		if node.Val <= prev {
			return false
		}
		prev = node.Val

		right := travese(node.Right)

		return left && right

	}

	return travese(root)
}

// @lc code=end
