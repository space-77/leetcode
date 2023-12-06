export default class ListNode {
    val;
    next;
    constructor(val = 0, next = null) {
        this.val = val;
        this.next = next;
    }
}
export function chain2Arr(node) {
    const arr = [];
    while (node) {
        arr.push(node.val);
        node = node.next;
    }
    return arr;
}
export function arr2Chain(arr) {
    const root = new ListNode();
    let node = root;
    arr.forEach((i) => {
        node.next = new ListNode(i);
        node = node.next;
    });
    return root.next;
}
