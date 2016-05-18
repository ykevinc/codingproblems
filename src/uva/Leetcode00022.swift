class Solution {
    func gen(l : Int, _ r : Int, _ s : String, inout _ o : [String]) {
                
        if l == 0 && r == 0 {
            o.append(s)
        }
        if l > 0 {
            gen(l - 1, r, s + "(", &o)
        }
        if r > 0 && r > l {
            gen(l, r - 1, s + ")", &o)
        }
    }
    func generateParenthesis(n: Int) -> [String] {
        var o = [String]()
        gen(n, n, "", &o)
        return o
    }
}
