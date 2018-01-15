package main

func longestSubstring(s string, k int) int {
	counts := make(map[rune]int, len(s))
	for _, c := range s {
		counts[c]++
	}
	splits := make([]int, 0, len(s))
	splits = append(splits, -1)
	for i, c := range s {
		if counts[c] < k {
			splits = append(splits, i)
		}
	}
	if len(splits) == 1 {
		return len(s)
	}
	splits = append(splits, len(s))
	max := 0
	for i := 1; i < len(splits); i++ {
		v := longestSubstring(s[splits[i-1]+1:splits[i]], k)
		if v > max {
			max = v
		}
	}
	return max
}
