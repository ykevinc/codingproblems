package main

func checkPossibility(nums []int) bool {
	mod := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if !mod {
				mod = true
				if i > 0 && nums[i-1] > nums[i+1] {
					nums[i+1] = nums[i]
				}
			} else {
				return false
			}
		}
	}
	return true
}
