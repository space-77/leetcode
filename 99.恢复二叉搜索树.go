package main

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
	nodeList := []*TreeNode{}
	var travese func(node *TreeNode)
	travese = func(node *TreeNode) {
		if node == nil {
			return
		}
		travese(node.Left)
		nodeList = append(nodeList, node)
		travese(node.Right)
	}
	travese(root)

	var first *TreeNode = nil
	var second *TreeNode = nil

	for i := 0; i < len(nodeList)-1; i++ {
		if nodeList[i].Val > nodeList[i+1].Val {
			if first == nil {
				first = nodeList[i]
			}
			second = nodeList[i+1]
		}
	}

	val := first.Val
	first.Val = second.Val
	second.Val = val
}

// @lc code=end
