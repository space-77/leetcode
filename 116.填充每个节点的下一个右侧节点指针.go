package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {

		levelLen := len(queue)
		for levelLen > 0 {

			levelLen--
			node := queue[0]
			queue = append(queue[:0], queue[1:]...)

			if levelLen > 0 {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}

	return root
}

// @lc code=end
