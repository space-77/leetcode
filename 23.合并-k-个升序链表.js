import ListNode, { chain2Arr, arr2Chain } from "./ListNode.js";

/*
 * @lc app=leetcode.cn id=23 lang=javascript
 *
 * [23] 合并 K 个升序链表
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
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function (lists) {
  //把所有节点取出放进一个数组
  let arr = lists.reduce((p, c) => {
    while (c) {
      p.push(c);
      c = c.next;
    }
    return p;
  }, []);
  let newHead = arr
    .sort((a, b) => a.val - b.val)
    .reduceRight((p, c) => {
      c.next = p;
      return c;
    }, null);
  return newHead;
};
// var mergeKLists = function (lists) {
//   const root = new ListNode();
//   let node = root;

//   lists = lists.filter(Boolean);

//   while (lists.length) {
//     const { min, index } = lists.reduce(
//       (pre, node, index) => {
//         if (pre.min.val <= node.val) {
//           return pre;
//         } else {
//           return { min: node, index };
//         }
//       },
//       { min: { val: Infinity }, index: -1 }
//     );

//     node.next = min;
//     node = node.next;

//     lists[index] = min.next;
//     lists = lists.filter(Boolean);
//   }

//   return root.next;
// };
// @lc code=end

const arr = [
  [1, 4, 5],
  [1, 3, 4],
  [2, 6],
];
const nodes = arr.map((i) => arr2Chain(i));
console.log(chain2Arr(mergeKLists(nodes)));
