package main

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}

	var travese func(node *TreeNode, num int, arr []int)
	travese = func(node *TreeNode, num int, arr []int) {
		if node == nil {
			return
		}

		val := node.Val
		arr = append(arr, val)

		if val == num && node.Left == nil && node.Right == nil {
			r := make([]int, len(arr))
			copy(r, arr)
			ans = append(ans, r)
		}

		nextNum := num - val
		travese(node.Left, nextNum, arr)
		travese(node.Right, nextNum, arr)

	}

	travese(root, targetSum, []int{})

	return ans
}

// @lc code=end
