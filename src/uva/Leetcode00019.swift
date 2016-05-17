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
    func count(head: ListNode?) -> Int {
        return head != nil ? (1 + count(head!.next)) : 0
    }
    func removeNthFromEnd(head: ListNode?, _ n: Int) -> ListNode? {
        let c = count(head)
        var l = c - n - 1
        if l == -1 {
            let newhead = head?.next
            head?.next = nil
            return newhead
        } else {
            var node = head
            for _ in 0.stride(to:l,by:1) {
                node = node!.next
            }
            node!.next = node!.next!.next
        }
        return head
    }
}
