package main

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tables := []struct {
		i []int
		o string
	}{
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{1}, "1"},
		{[]int{1, 2, 2, 1}, "2211"},
		{[]int{0, 0}, "0"},
	}
	for _, table := range tables {
		v := largestNumber(table.i)
		if v != table.o {
			t.Errorf("f(%v) != %s, but %s", table.i, table.o, v)
		}
	}
}
