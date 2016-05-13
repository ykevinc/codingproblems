class Solution {
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
            let r = d - Int(pow(10.0, Double(l - 1)))*(d%10)
            if r == 0 {
                return true
            } else if r < 0 {
                return false
            } else if r >= Int(pow(10.0, Double(l - 1))) {
                return false
            }
            d = r / 10
            l -= 2
        }
        return true
    }
}
