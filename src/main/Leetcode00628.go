package main

import (
	"sort"
)

func maximumProduct(nums []int) int {
	// all postive
	// two negative, one postive
	sort.Ints(nums)
	v1 := nums[0] * nums[1] * nums[len(nums)-1]
	v2 := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	if v1 > v2 {
		return v1;
	} else if v2 > v1 {
		return v2;
	} else {
		return v1;
	}
}
