class Solution {
    func convert(s: String, _ numRows: Int) -> String {
        var arr = Array(count: numRows, repeatedValue: Array(count: min(s.characters.count, 500), repeatedValue: Character("\0")))
        var x = 0, y = 0, dir = 0
        var dirs : [[Int]] = [[1,0],[-1,1],[0,1]]
        var maxy = 0
        for c in s.characters {
            arr[x][y] = c
            if numRows == 1 {
                dir = 2
            } else if x == numRows - 1 {
                dir = 1
            } else if x == 0 {
                dir = 0
            }
            x += dirs[dir][0]
            y += dirs[dir][1]
            maxy = max(maxy, y)
        }
        var sb = ""
        var array = [String]()
        for x in 0..<numRows {
            for y in 0..<maxy {
                if arr[x][y] != "\0" {
                    array.append(String(arr[x][y]))
                } 
            }
        }
        return array.joinWithSeparator("")
    }
}
