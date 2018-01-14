package main

func dominantIndex(nums []int) int {
	m := make([]int, 100)
	maxI := -1
	for i, n := range nums {
		m[n] += 1
		if maxI == -1 || n > nums[maxI] {
			maxI = i
		}
	}
	if maxI != -1 {
		maxV := nums[maxI]
		m[maxV] -= 1
		for v := maxV; v >= 0; v-- {
			if m[v] > 0 {
				if v*2 > maxV {
					return -1
				} else {
					return maxI
				}
			}
		}
	}
	return maxI
}
