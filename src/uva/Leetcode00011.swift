class Solution {
    func maxArea(height: [Int]) -> Int {
        var l = 0,  r = height.count-1, maxArea = 0
        while l <= r {
            var area = min(height[l], height[r])*(r-l) 
            maxArea = max(maxArea, area)
            if height[l] > height[r] {
                r -= 1
            } else if height[r] > height[l] {
                l += 1
            } else {
                l += 1
            }
        }
        return maxArea
    }
}
