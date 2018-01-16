package main

import (
	"strings"
	"strconv"
)

func largestNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}
	m := make([][]string, 10, 10)
	for _, s := range strs {
		m[s[0]-'0'] = append(m[s[0]-'0'], s)
	}
	if len(m[0]) == len(nums) {
		return "0"
	}
	max := ""
	dfs(m, 9, "", &max)
	return max
}

func dfs(m [][]string, d int, c string, max *string) {
	if d == -1 {
		return
	}
	if len(c) > len(*max) {
		*max = c
	} else if len(c) == len(*max) {
		for i := 0; i < len(c); i++ {
			if c[i] > (*max)[i] {
				*max = c
				break
			} else if c[i] < (*max)[i] {
				break
			}
		}
	}
	v := m[d]
	choices := 0
	for i, s := range v {
		if s == "@" {
			continue
		}
		choices++
		v[i] = "@"
		dfs(m, d, c+s, max)
		v[i] = s
	}
	if choices == 0 {
		dfs(m, d-1, c, max)	
	}
}
