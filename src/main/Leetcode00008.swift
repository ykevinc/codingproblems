extension String {
    func ascii() -> Int {
      var retVal : UInt32 = 0
      for val in self.unicodeScalars where val.isASCII() {
          retVal = UInt32(val)
      }
      return Int(retVal)
    }
    func trim() -> String
    {
        return self.stringByTrimmingCharactersInSet(NSCharacterSet.whitespaceCharacterSet())
    }
}

class Solution {
    var INVALID_POS  = 2147483647
    var INVALID_NEG  = -2147483648
    func returnInvalid(isPostive: Bool) -> Int {
        return isPostive ? INVALID_POS : INVALID_NEG
    }
    func myAtoi(str: String) -> Int {
        var letters = str.trim().characters.map{String($0)}, isPostive = true
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
        for c in letters {
            if c < "0" || c > "9" {
                break
            }
            (val, didOverflow) = Int.multiplyWithOverflow(val, 10)
            if didOverflow {
                return returnInvalid(isPostive)
            }
            (val, didOverflow) = Int.addWithOverflow(c.ascii()-"0".ascii(), val)
            if didOverflow {
                return returnInvalid(isPostive)
            }
        }
        if val > INVALID_POS {
            return returnInvalid(isPostive)
        }
        return isPostive ? Int(val) : -Int(val)
    }
}
