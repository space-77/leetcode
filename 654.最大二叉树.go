package main

/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
 */
// @lc code=start

func getMax(nums []int) (index int, val int) {
	index = 0
	val = nums[index]

	for i := 0; i < len(nums); i++ {
		if nums[i] > val {
			index = i
			val = nums[i]

		}
	}
	return index, val
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index, _ := getMax(nums)

	node := &TreeNode{
		nums[index],
		constructMaximumBinaryTree(nums[:index]),
		constructMaximumBinaryTree(nums[index+1:]),
	}

	return node

}

// @lc code=end
