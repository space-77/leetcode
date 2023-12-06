import ListNode, { arr2Chain, chain2Arr } from "./ListNode.js";

/*
 * @lc app=leetcode.cn id=25 lang=javascript
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var reverseKGroup = function (head, k) {
  const stack = [];
  const root = new ListNode();
  root.next = head;

  let pre = root;
  let index = 0;

  while (head) {
    
    index++;
    const done = index % k === 0;
    stack.push(head);
    head = head.next;
    
    if (done) {
      while (stack.length) {
        pre.next = stack.pop();
        pre = pre.next;
      }
    }

    pre.next = head || stack[0] || null
  }

  return root.next;
};
// @lc code=end

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];
const node = arr2Chain(arr);
// chain2Arr(reverseKGroup(node, 2));
console.log(chain2Arr(reverseKGroup(node, 2)));
