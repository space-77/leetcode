package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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

// func rightSideView(root *TreeNode) (ans []int) {

// 	var dfs func(node *TreeNode, depth int)

// 	dfs = func(node *TreeNode, depth int) {
// 		if node == nil {
// 			return
// 		}

// 		// ans中还无“depth深度”的值，而第一次遍历到depth深度，故加入
// 		if depth == len(ans) {
// 			ans = append(ans, node.Val)
// 		}
// 		dfs(node.Right, depth+1)
// 		dfs(node.Left, depth+1)

//		}
//		dfs(root, 0)
//		return
//	}
func rightSideView(root *TreeNode) (ans []int) {

	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		index := 0
		levelLen := len(queue)

		for index < levelLen {
			index++
			node := queue[0]

			if index == levelLen {
				ans = append(ans, node.Val)
			}

			queue = append(queue[:0], queue[1:]...)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
	}

	return
}

// @lc code=end
