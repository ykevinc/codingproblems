class Solution {
    func convert(s: String, _ numRows: Int) -> String {
        var letters = s.characters.map{Character(String($0))}
        var arr = Array(count: numRows, repeatedValue: Array(count: s.characters.count, repeatedValue: Character("\0")))
        var x = 0, y = 0
        var dirDown = true
        for c in letters {
            arr[x][y] = c
            if dirDown {
                (x, y, dirDown) = (x + 1) % numRows == 0 ? (max(x - 1, 0), y + 1, false) : (min(x + 1,numRows-1), y, true)
            } else {
                (x, y, dirDown) = x == 0 ? (min(x + 1,numRows-1), y, true) : (max(x - 1, 0), y + 1, false)
            }
        }
        var sb = ""
        for x in 0..<numRows {
            for y in 0..<s.characters.count {
                if arr[x][y] != "\0" {
                    sb += String(arr[x][y])
                } 
            }
        }
        return sb
    }
}
