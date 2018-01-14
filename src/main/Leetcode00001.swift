class Solution {
    func twoSum(nums: [Int], _ target: Int) -> [Int] {
        var sums : [Int : Int] = [:]
        var i = 0
        while i < nums.count {
            let n = nums[i]
            if let j = sums[target - n] {
                return [j, i]
            }
            sums[n] = i
            i = i + 1
        }
        return []
    }
}
