package main

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	deleted := make([]bool, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		} else if intervals[i].Start > intervals[j].Start {
			return false
		} else {
			return intervals[i].End < intervals[j].End
		}
	})

	for i, j := 0, 1; i < len(intervals)-1 && j < len(intervals); {
		if intervals[j].Start <= intervals[i].End {
			if intervals[j].End > intervals[i].End {
				intervals[i].End = intervals[j].End
				deleted[j] = true
			} else {
				deleted[j] = true
			}
			j++
		} else {
			i = j
			j = i + 1
		}
	}
	i := 0
	for j := i; i < len(intervals); i++ {
		for j < len(intervals) && deleted[j] {
			j++
		}
		if j == len(intervals) {
			break
		}
		if i != j {
			intervals[i] = intervals[j]
			deleted[j] = true
		}
		j = i + 1
	}
	return intervals[0:i]
}

func mergeInts(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] <= intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	for i, j := 0, 1; i < len(intervals)-1 && j < len(intervals); {
		//fmt.Println(i, j, len(intervals), len(intervals[j]), len(intervals[i]))
		//fmt.Println(intervals)
		if intervals[j][0] < intervals[i][1] {
			if intervals[j][1] > intervals[i][1] {
				intervals[i][1] = intervals[j][1]
				intervals[j] = nil
			} else {
				intervals[j] = nil
			}
			j++
		} else {
			i = j
			j = i + 1
		}
	}
	i := 0
	for j := i; i < len(intervals); i++ {
		for j < len(intervals) && intervals[j] == nil {
			j++
		}
		if j == len(intervals) {
			break
		}
		if i != j {
			intervals[i] = intervals[j]
			intervals[j] = nil
		}
		j = i + 1
	}
	return intervals[0:i]
}
