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
 
 class Heap <T:IntegerType> {
    var items = [T](count:1, repeatedValue:0)
    var n = 0
    func push(t : T) {
        items.append(t)
        n += 1
        swim(n)
    }
    func pop() -> T? {
        if n == 0 {
            return nil
        } else if n > 1 {
            swap(&items[1], &items[n])
        }
        n -= 1
        sink(1)
        return items.removeLast()
    }
    func count() -> Int {
        return n
    }
    func sink(index : Int) {
        var i = index
        while i*2 <= n {
            var j = i*2
            if j + 1 <= n && items[j+1] < items[j] {
                j += 1
            }
            if items[i] < items[j] {
                break
            }
            swap(&items[i], &items[j])
            i = j
        }
        
    }
    func swim(index : Int) {
        var i = index
        while(i > 1 && items[i] < items[i/2]) {
            swap(&items[i], &items[i/2])
            i = i/2
        }
    }
}

class Solution {
    func mergeKLists(lists: [ListNode?]) -> ListNode? {
        var heap = Heap<Int>()
        for list in lists {
            var node = list
            while node != nil {
                print(node!.val)
                heap.push(node!.val)
                node = node!.next
            }
        }
        var prehead = ListNode(0), node : ListNode? = prehead
        while let val = heap.pop() {
            print(val)
            node!.next = ListNode(val)
            node = node!.next
        } 
        return prehead.next
    }
}
