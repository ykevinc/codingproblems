package main

import (
	"testing"
)

func TestMaximumProduct(t *testing.T) {
	tables := []struct {
		i []int
		o int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-10, -20, 3, 1}, 600},
	}
	for _, table := range tables {
		v := maximumProduct(table.i)
		if v != table.o {
			t.Errorf("f(%v) != %d, but %d", table.i, table.o, v)
		}
	}
}
