class Solution {
    func findMedianSortedArrays(nums1: [Int], _ nums2: [Int]) -> Double {
        var n = 0, i = 0, j = 0
        var m = nums1.count + nums2.count
        var a = Array<Int>(count: m, repeatedValue: 0)
        //even (m-1)/2 m/2
        //odd m/2
        while i < nums1.count || j < nums2.count {
            if i < nums1.count && j < nums2.count {
                if nums1[i] > nums2[j] {
                    a[n] = nums2[j]
                    j = j + 1
                } else if nums1[i] < nums2[j] {
                    a[n] = nums1[i]
                    i = i + 1
                } else {
                    a[n] = nums1[i]
                    i = i + 1
                }
            } else if i < nums1.count {
                a[n] = nums1[i]
                i = i + 1
            } else if j < nums2.count {
                a[n] = nums2[j]
                j = j + 1
            }
            if n >= m/2 {
                if m % 2 == 0 {
                    return Double(a[(m-1)/2] + a[m/2])/2.0
                } else {
                    return Double(a[m/2])
                }
            }
            n = n + 1
        }
        return -1.0
    }
}
