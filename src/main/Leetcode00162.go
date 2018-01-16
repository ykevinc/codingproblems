package main

func findPeakElement(nums []int) int {
	l := len(nums)
	if l >= 2 {
		if nums[0] > nums[1] {
			return 0
		}
		if nums[l-1] > nums[l-2] {
			return l-1
		}
		for i := 1;i < l-1;i++ {
			if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return 0
}
