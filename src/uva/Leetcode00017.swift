class Solution {
    func combo(digits : [String], _ pad : [String:String], _ i : Int, _ s : String, inout _ c : [String]) {
        if i == digits.count {
            c.append(s)
            return
        }
        for p in pad[digits[i]]!.characters {
            combo(digits, pad, i+1, s + String(p), &c)
        }
    }

    func letterCombinations(digits: String) -> [String] {
        var pad = [
            "2":"abc",
            "3":"def",
            "4":"ghi",
            "5":"jkl",
            "6":"mno",
            "7":"pqrs",
            "8":"tvu",
            "9":"wxyz",
        ]
        if digits.characters.count == 0 {
            return []
        }
        var c = [String]()
        combo(digits.characters.map{String($0)}, pad, 0, "", &c)
        return c
    }
}
