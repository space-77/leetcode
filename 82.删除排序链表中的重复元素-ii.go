package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	root := &ListNode{Next: head}
	preNode := root

	for head != nil {
		next := head.Next

		if next != nil && next.Val == head.Val {
			for next != nil && next.Val == head.Val {
				next = next.Next
			}
			preNode.Next = next
		} else {
			preNode = head
		}

		head = next
	}

	return root.Next
}

// @lc code=end
