class Solution {
    func romanToInt(s: String) -> Int {
        let letters = s.characters.map{String($0)}
        let r2n = [
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
        ]
        var result = 0;
        for i in 0..<letters.count {
           if(i+1<letters.count&&r2n[letters[i]]<r2n[letters[i+1]]) {
               result -= r2n[letters[i]]!;
           } else {
               result += r2n[letters[i]]!;
           }
        }
        return result
    }
}
