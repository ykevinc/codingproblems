class Solution {
    func intToRoman(num: Int) -> String {
        let romans = ["I", "IV", "V", "IX","X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"]
        let nums = [1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000]
        var val = num, s = "", i = 12
        while i >= 0 {
            if val>=nums[i] {
                s += romans[i];
                val -= nums[i];
            } else {
                i--;
            }
        }
        return s
    }
}
