package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}

	sort.Slice(strs, func(i, j int) bool {
		s1 := strs[i] + strs[j]
		s2 := strs[j] + strs[i]
		return s1 > s2
	})

	if strs[0] == "0" {
		return "0"
	}

	return strings.Join(strs, "")
}
