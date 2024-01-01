/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	fast := head
	slwo := head

	for fast != nil && fast.Next != nil {
		slwo = slwo.Next
		fast = fast.Next.Next
	}

	return slwo

}

// @lc code=end

