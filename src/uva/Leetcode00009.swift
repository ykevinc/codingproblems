class Solution {
    func getTens(p: Int) -> Int {
        var s = 1
        for i in 0..<p {
            s *= 10
        }
        return s
    }
    func isPalindrome(x: Int) -> Bool {
        var l = 0
        if x < 0 {
            return false
        }
        var d = abs(x)
        while d > 0 {
            d /= 10
            l += 1
        }
        d = abs(x)
        while d > 0 && l > 1 {
            let tens = getTens(l-1), r = d-tens*(d % 10)
            if r == 0 {
                return true
            } else if r < 0 {
                return false
            } else if r >= tens {
                return false
            }
            d = r / 10
            l -= 2
        }
        return true
    }
}
