package main

import "math"

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	var first *TreeNode
	var second *TreeNode
	var preNode = &TreeNode{Val: math.MinInt}

	var travese func(node *TreeNode)
	travese = func(node *TreeNode) {
		if node == nil {
			return
		}
		travese(node.Left)

		if preNode.Val > node.Val {
			if first == nil {
				first = preNode
			}
			second = node
		}

		preNode = node

		travese(node.Right)
	}
	travese(root)

	first.Val, second.Val = second.Val, first.Val
}

// @lc code=end
