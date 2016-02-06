/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Leetcode00019 {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode node = head;
        int length = 0;
        while (node != null) {
            node = node.next;
            length++;
        }
        int cut = length - n - 1;
        if (cut < 0) {
            return head.next;
        }
        node = head;
        while (node.next != null && cut-- > 0) {
            node = node.next;
        }
        node.next = node.next.next;
        return head;
    }
}
