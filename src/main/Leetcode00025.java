/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Leetcode00025 {
    public ListNode reverseKGroup(ListNode head, int k) {
       if (head == null) {
           return null;
       }
       ListNode l = head, r, previous = null, next = null, node = null, newHead, rnext, lprevious=null;
       r = newHead = getKth(head, k-1);
       if (r == null) {
           return head;
       }
       for (next=(head.next!=null?head.next:null),l = head;r!=null;r=getKth(l, k-1)) {
           if (lprevious != null) {
               lprevious.next = r;
           }
           rnext = r.next;
           lprevious = l;
           while(l != rnext) {
               next = l.next;
               l.next = previous;
               previous = l;
               l = next;
           }
           lprevious.next = l = rnext;
       }
       return newHead;
    }
    
    public ListNode getKth(ListNode h, int k) {
        if (h == null) {
            return null;
        } else if (k == 0) {
            return h;
        } else {
            return getKth(h.next, k-1);
        }
    }
    
}
