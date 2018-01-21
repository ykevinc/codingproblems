package main

import (
	"testing"
)

func TestSmallestDistancePair(t *testing.T) {
	tables := []struct{
		i []int
		k int
		o int
	} {
		{[]int{1,3,1}, 1, 0},
	}
	for _, table := range tables {
		v := smallestDistancePair(table.i, table.k)
		if v != table.o {
			t.Errorf("f(%v, %d) != %d, but %d", table.i, table.k, table.o, v)
		}
	}
}