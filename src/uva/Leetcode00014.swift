class Solution {
    func longestCommonPrefix(strs: [String]) -> String {
        var lettersOfArray = strs.map{$0.characters.map{String($0)}}
        let minLength = strs.count > 0 ? strs.reduce(Int.max, combine:{min($0, $1.characters.count)}) : 0
        var sb = ""
        for i in 0..<minLength {
            var c = "\0"
            if lettersOfArray.count > 0 {
                c = lettersOfArray[0][i]
            }
            for j in 1..<lettersOfArray.count {
                if lettersOfArray[j][i] != c {
                    return sb
                }
            }
            sb += String(c)                
        }
        return sb
    }
}
