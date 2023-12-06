export default class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val?: number, next?: ListNode | null);
}
export declare function chain2Arr(node: ListNode | null): number[];
export declare function arr2Chain(arr: number[]): ListNode;
