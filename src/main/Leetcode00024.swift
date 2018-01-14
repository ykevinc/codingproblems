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
    func swapPairs(head: ListNode?) -> ListNode? {
        var first : ListNode? = head
        var second : ListNode? = head?.next ?? nil
        var third : ListNode?
        var prehead = ListNode(0), previous : ListNode? = prehead
        if second == nil {
            return first
        }
        while first != nil && second != nil {
            third = second?.next
            
            first!.next = third
            second!.next = first
            previous?.next = second
            
            previous = first
            first = third
            second = third?.next
            
        }
        return prehead.next
    }
}
