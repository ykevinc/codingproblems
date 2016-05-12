class Solution {
    func longestPalindrome(s: String) -> String {
        let a = s.characters.map { String($0) }
        var max = 0, maxL = 0, maxR = 0
        for i in 0.stride(through: a.count, by: 1)  {
            var (length, l, r) = getLongestMatch(a, i-1, i+1)
            if length > max {
                (max, maxL, maxR) = (length, l, r)
            }
            (length, l, r) = getLongestMatch(a, i, i+1)
            if length > max {
                (max, maxL, maxR) = (length, l, r)
            }
        }
        var f = ""
        for i in maxL...maxR {
            f += a[i] 
        }
        return f
    }
    func getLongestMatch(a : Array<String>, _ l : Int, _ r : Int) -> (length:Int,l:Int,j:Int) {
        var length = 0, index = 0, i = l, j = r
        while i >= 0 && i < a.count && j >= 0 && j < a.count {
            if a[i] != a[j] {
                break
            } else {
                length = max(length, j - i + 1)
            }
            i = i - 1
            j = j + 1
        }
        return (length, i + 1, j - 1)
    }
}
