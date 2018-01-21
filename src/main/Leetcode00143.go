package main

type ListNode struct {
	 Val int
	 Next *ListNode
}
 
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
    m := head
    move := true
    for p := head.Next;p.Next != nil;p = p.Next {
        if move {
            m = m.Next            
        }
        move = !move
    }
    left := m.Next
    right := left.Next
    if right != nil {
        for right.Next != nil {
            c := right.Next
            right.Next = c.Next
            c.Next = left.Next
            left.Next = c
        }        
    }
    p := head
    p2 := left.Next
    for p2 != nil {
        left.Next = p2.Next
        p2.Next = p.Next
        p.Next = p2
        p = p2.Next
        p2 = left.Next
    }
    p2 = p.Next
    p.Next = left
    left.Next = p2
    p2.Next = nil
}