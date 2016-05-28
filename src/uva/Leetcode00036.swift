extension String {
    func ascii() -> Int {
      var retVal : UInt32 = 0
      for val in self.unicodeScalars where val.isASCII() {
          retVal = UInt32(val)
      }
      return Int(retVal)
    }
}

class Solution {
    func isValidSudoku(board: [[Character]]) -> Bool {
        var xx = Array(count:9, repeatedValue:Array(count:10, repeatedValue:false))
        var yy = Array(count:9, repeatedValue:Array(count:10, repeatedValue:false))
        var xy = Array(count:9, repeatedValue:Array(count:10, repeatedValue:false))
        for x in 0..<9 {
            for y in 0..<9 {
                if board[x][y] != "." {
                    let ascii = String(board[x][y]).ascii() - "0".ascii()
                    if xx[x][ascii] {
                        return false
                    }
                    xx[x][ascii] = true
                    if yy[y][ascii] {
                        return false
                    }
                    yy[y][ascii] = true
                    if xy[x/3*3+y/3][ascii] {
                        return false
                    }
                    xy[x/3*3+y/3][ascii] = true
                }
                

            }
        }
        return true
    }
}
