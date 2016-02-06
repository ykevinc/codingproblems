/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Leetcode00024 {
    public ListNode swapPairs(ListNode head) {
        if (head == null) {
            return null;
        }
        if (head.next != null) {
            ListNode second = head.next;
            head.next = swapPairs(head.next.next);
            second.next = head;
            return second;
        } else {
            return head;            
        }
    }
}
