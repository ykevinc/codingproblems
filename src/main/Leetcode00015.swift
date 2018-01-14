class Solution {
    func threeSum(nums: [Int]) -> [[Int]] {
        var m = [String : [Int]]()
        let arr = nums.sort()
        for i in 0.stride(to: arr.count-2, by: 1) {
            var j = i + 1
            var k = arr.count - 1
            while j < k {
                let sum = arr[i] + arr[j] + arr[k]
                if sum == 0 {
                    let key = "(\(arr[i]), \(arr[j]), \(arr[k]))"
                    if m[key] == nil {
                        m[key] = [arr[i], arr[j], arr[k]]
                    }
                    k -= 1
                    j += 1
                } else if sum > 0 {
                    k -= 1
                } else {
                    j += 1
                }
            }
        }
        return Array(m.values)
    }
}
