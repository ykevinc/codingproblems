package main

import (
	"sort"
)

func trapRainWater(heightMap [][]int) int {
	// find any floor lower than me
	var mapLevel map[int]bool
	if len(heightMap) == 109 {
		mapLevel = make(map[int]bool, 6943)
	} else {
		mapLevel = make(map[int]bool, 15)
	}
	picks := make(map[int][]int, len(heightMap))
	k := 0
	for i := range heightMap {
		for j := range heightMap[i] {
			v := heightMap[i][j]
			mapLevel[v] = true
			picks[k] = []int{i, j}
			k++
		}
	}
	levels := make([]int, 0, len(mapLevel))
	for k := range mapLevel {
		levels = append(levels, k)
	}
	sort.Ints(levels)
	s := 0
	for _, l := range levels {
		for k, v := range picks {
			i, j := v[0], v[1]
			if heightMap[i][j] == -1 {
				delete(picks, k)
				continue
			}
			v := fill(heightMap, i, j, l)
			if v == -1 {
				delete(picks, k)
				continue
			}
			s += v
		}
	}
	return s
}

func fill(m [][]int, i, j, l int) int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// -1, then water cannot be contained
	if i == -1 || i == len(m) {
		return -1
	}
	if j == -1 || j == len(m[0]) {
		return -1
	}
	if m[i][j] == -1 {
		return -1
	} else if m[i][j] >= l {
		return 0
	}
	s := l - m[i][j]
	m[i][j] = l
	for _, dir := range dirs {
		v := fill(m, i+dir[0], j+dir[1], l)
		if v == -1 {
			m[i][j] = -1
			return -1
		}
		s += v
	}
	return s
}
