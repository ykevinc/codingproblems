package main

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	tables := []struct{
		i int
		o []int
	}{
		{3, []int{0,1,1,2}},
		{5, []int{0,1,1,2,1,2}},
		{8, []int{0,1,1,2,1,2,2,3,1}},
	}
	for _, table := range tables {
		v := countBits(table.i)
		if !intsEquals(v, table.o) {
			t.Errorf("f(%d) != %v, but %v", table.i, table.o, v)	
		}
	}
}

func intsEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}