class Solution {
    func removeElement(inout nums: [Int], _ val: Int) -> Int {
        var l = 0
        var r = 0
        while r < nums.count {
            if nums[r] != val {
                if l != r {
                    swap(&nums[l], &nums[r])
                }
                l += 1
            }
            r += 1
        }
        return l
    }
}
