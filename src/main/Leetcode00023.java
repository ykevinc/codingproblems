import java.util.*;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Leetcode00023 {
    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> p = new PriorityQueue<ListNode>(10, new Comparator<ListNode>()
            {
                public int compare(ListNode o1, ListNode o2) {
                    return o1.val - o2.val;
                }
            }
        );
        for (ListNode l : lists) {
            if (l != null) {
                p.add(l);
            }
        }
        ListNode prehead = new ListNode(0);
        ListNode node = prehead;
        while (!p.isEmpty()) {
            ListNode next = p.poll();
            if (next.next != null) {
                p.add(next.next);
            }
            node.next = next;
            node = node.next;
        }
        return prehead.next;
    }
}
