package main

import (
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tables := []struct {
		i []int
		o int
	}{
		{[]int{1, 2, 3, 1}, 2},
	}
	for _, table := range tables {
		v := findPeakElement(table.i)
		if v != table.o {
			t.Errorf("f(%v) != %d", table.i, table.o)
		}
	}
}
