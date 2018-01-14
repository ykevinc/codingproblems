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
    func reverse(i: Int) -> Int {
        if i == 0 {
            return 0
        }
        var letters = String(i).characters.map{String($0)}, isPostive = true
        if letters.count == 0 || letters.count == 1 && (letters[0] == "+" || letters[0] == "-")  {
            return 0
        }
        if letters.count > 2 {
            if letters[0] == "+" && letters[1] == "-" || letters[0] == "-" && letters[1] == "+" {
                return 0
            }
        }
        if letters[0] == "-" {
            isPostive = false
            letters = Array(letters[1..<letters.count])
        } else if letters[0] == "+" {
            letters = Array(letters[1..<letters.count])
        }
       
        var val = 0, didOverflow = false
        for c in letters.reverse() {
            if c < "0" || c > "9" {
                break
            }
            (val, didOverflow) = Int.multiplyWithOverflow(val, 10)
            if didOverflow {
                return 0
            }
            (val, didOverflow) = Int.addWithOverflow(c.ascii()-"0".ascii(), val)
            if didOverflow {
                return 0
            }
        }
        if val > 2147483647 {
            return 0
        }
        return isPostive ? Int(val) : -Int(val)
    }
}
