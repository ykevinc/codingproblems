package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j := 0, 0
	m := 0
	w := make(map[byte]bool, len(s))
	for i < len(s) && j < len(s) {
		c := s[j]
		if _, ok := w[c]; !ok {
			w[c] = true
			if len(w) > m {
				m = len(w)
			}
			j++
		} else {
			delete(w, s[i])
			i++
		}
	}
	return m
}
