package main

import "testing"

func Test(t *testing.T) {
	tables := []struct {
		i []int
		o int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}
	for _, table := range tables {
		v := dominantIndex(table.i)
		if v != table.o {
			t.Errorf("f(%v) != %d but %d\n", table.i, table.o, v)
		}
	}
}
