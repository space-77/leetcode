package main

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	root := &ListNode{Next: head}
	head = root
	index := 0
	var preNode *ListNode

	for head != nil {
		if index == left {

			var pre *ListNode
			last := head
			for index <= right {
				next := head.Next
				head.Next = pre
				pre = head
				head = next
				index++
			}
			last.Next = head
			preNode.Next = pre
			break
		} else {
			preNode = head
			head = head.Next
			index++
		}
	}

	return root.Next
}

// @lc code=end
