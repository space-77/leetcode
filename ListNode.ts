export default class ListNode {
  constructor(public val = 0, public next: ListNode | null = null) {}
}

export function chain2Arr(node: ListNode | null) {
  const arr: number[] = [];
  while (node) {
    arr.push(node.val);
    node = node.next;
  }

  return arr;
}

export function arr2Chain(arr: number[]): ListNode {
  const root = new ListNode();

  let node = root;

  arr.forEach((i) => {
    node.next = new ListNode(i);
    node = node.next;
  });

  return root.next!;
}
