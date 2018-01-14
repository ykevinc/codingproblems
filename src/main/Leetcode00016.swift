class Solution {
    func threeSumClosest(nums: [Int], _ target: Int) -> Int {
        var ii : Int?, jj : Int?, kk :Int?
        var closest = Int.max
        let arr = nums.sort()
        for i in 0.stride(to: arr.count-2, by: 1) {
            var j = i + 1
            var k = arr.count - 1
            while j < k {
                let sum = arr[i] + arr[j] + arr[k]
                if ii == nil || abs(sum - target) < abs(closest - target)   {
                    ii = i
                    jj = j
                    kk = k
                    closest = sum
                }
                if sum == target {
                    return target
                    k -= 1
                    j += 1
                } else if sum > target {
                    k -= 1
                } else {
                    j += 1
                }
            }
        }
        return closest
    }
}
