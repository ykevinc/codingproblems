class Solution {
    func fourSum(nums: [Int], _ target: Int) -> [[Int]] {
        var m = [String : [Int]]()
        let arr = nums.sort()
        for i in 0.stride(to: arr.count-3, by: 1) {
            for j in (i+1).stride(to: arr.count-2, by: 1) {
                var k = j + 1
                var l = arr.count - 1
                while k < l {
                    let sum = arr[i] + arr[j] + arr[k] + arr[l]
                    if sum == target {
                        let key = "(\(arr[i]), \(arr[j]), \(arr[k]), \(arr[l]))"
                        if m[key] == nil {
                            m[key] = [arr[i], arr[j], arr[k], arr[l]]
                        }
                        l -= 1
                        k += 1
                    } else if sum > target {
                        l -= 1
                    } else {
                        k += 1
                    }
                }
            }

        }
        return Array(m.values)
    }
}
