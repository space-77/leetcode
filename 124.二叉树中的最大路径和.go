package main

import "math"

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	max := math.MinInt

	var travese func(node *TreeNode) int
	travese = func(node *TreeNode) int {

	}

	travese(root)
	return max

}

// @lc code=end
