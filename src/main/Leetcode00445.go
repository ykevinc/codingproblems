package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1 = reverseList(l1)
    l2 = reverseList(l2)
    c := 0
    pH := &ListNode{0, nil}
    n := pH
    for l1 != nil || l2 != nil || c != 0 {
        v := 0
        if l1 != nil {
            v += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v += l2.Val
            l2 = l2.Next
        }
        v += c  
        c = 0
        if v > 9 {
            c = v/10
            v %= 10
        }
        n.Next = &ListNode{v,  nil}
        n = n.Next
    }
    return reverseList(pH.Next)
}

func reverseList(h *ListNode) *ListNode {
    var newH *ListNode
    for h != nil {
        next := h.Next
        h.Next = newH
        newH = h
        h = next
    }
    return newH
}