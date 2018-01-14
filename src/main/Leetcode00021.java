/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Leetcode00021 {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode prehead = new ListNode(0);
        ListNode node = prehead;
        while (l1 != null || l2 != null) {
            ListNode choice;
            if (l1 != null && l2 != null) {
                if (l1.val <= l2.val) {
                    choice = l1;
                    l1 = l1.next;
                } else {
                    choice = l2;
                    l2 = l2.next;             
                }
            } else if (l1 != null) {
                choice = l1;
                l1 = l1.next; 
            } else {
                choice = l2;
                l2 = l2.next;      
            }
            node.next = choice;
            node = choice;
        }
        return prehead.next;
    }
}
