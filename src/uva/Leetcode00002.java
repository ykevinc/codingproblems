/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = null;
        ListNode node = null;
        ListNode l = l1,r = l2;
        int carry = 0;
        if (l != null||r != null) {
            int sum = sum(l,r,carry);
            carry = sum/10;
            sum = sum%10;
            head = node = new ListNode(sum);
            l = (l!=null?l.next:null);
            r = (r!=null?r.next:null);
            while (l!=null||r!=null||carry!=0) {
                sum = sum(l,r,carry);
                carry = sum/10;
                sum = sum%10;
                node.next = new ListNode(sum);
                node = node.next;
                l = (l!=null?l.next:null);
                r = (r!=null?r.next:null);
            }
        }
        return head;
    }
    public int sum(ListNode l, ListNode r, int carry) {
        int sum  = 0;
        if (l != null) {
            sum += l.val;
        }
        if (r != null) {
            sum += r.val;
        }
        sum += carry;
        return sum;
    }
}
