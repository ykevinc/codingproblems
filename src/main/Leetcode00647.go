package main

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	c := 0
	for i := 0; i < len(s); i++ {
		c += countSubstringsSub(s, i, i) + countSubstringsSub(s, i, i+1)
	}
	return c
}

func countSubstringsSub(s string, i, j int) int {
	c := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		c++
	}
	return c
}
