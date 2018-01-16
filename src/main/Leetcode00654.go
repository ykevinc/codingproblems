package main

func findErrorNums(nums []int) []int {
	m := make(map[int]int, len(nums))
	k := 0
	for _, n := range nums {
		m[n]++
		if m[n] == 2 {
			k = n
		}
	}
	for i := 1;i <= len(nums);i++ {
		if m[i] == 0 {
			return []int{k, i}
		}
	}
	return nil
}
