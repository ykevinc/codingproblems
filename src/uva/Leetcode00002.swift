/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init(_ val: Int) {
 *         self.val = val
 *         self.next = nil
 *     }
 * }
 */
class Solution {
    func addTwoNumbers(l1: ListNode?, _ l2: ListNode?) -> ListNode? {
        var l : ListNode? = l1, r : ListNode? = l2, c = 0
        let head = ListNode(0)
        var node : ListNode = head;
        var vv = l != nil ? 1 : 0
        while l != nil || r != nil || c > 0 {
            var s : Int = (l?.val ?? 0) + (r?.val ?? 0) + c
            c = s / 10
            s = s % 10
            node.next = ListNode(s)
            node = node.next!
            l = l?.next
            r = r?.next
        }
        return head.next
    }
}
