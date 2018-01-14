extension String {
    subscript (i: Int) -> Character {
        return self[self.startIndex.advancedBy(i)]
    }
    subscript (r: Range<Int>) -> String {
        return String(Array(arrayLiteral:self)[r.startIndex...r.endIndex])
    }
    var ascii: Int {
        return Int(self.unicodeScalars.first!.value)
    }
}
class Solution {
    func lengthOfLongestSubstring(s: String) -> Int {
        var mapIndexByChar = Array<Int?>(count: 256, repeatedValue: nil), i = 0, j = 0, l = 0
        let letters = s.characters.map { String($0) }
        while i < letters.count {
            var v = letters[i].ascii
            if let index = mapIndexByChar[v] {
                 j = max(j, index + 1)
            }
            mapIndexByChar[v] = i
            l = max(l, i - j + 1)
            i = i + 1
        }
        return l
    }
}
