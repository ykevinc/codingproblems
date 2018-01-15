package main

import (
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	graph1 := [][]int{
		{1, 1, 10},
		{1, 2, 10},
		{2, 3, 10},
		{2, 3, 10},
	}
	graph2 := [][]int{
		{1, 2, 1},
	}
	graph3 := [][]int{
		{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2},
	}
	graph4 := [][]int{
		{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74},
	}
	tables := []struct {
		i [][]int
		n int
		k int
		o int
	}{
		{graph1, 3, 1, 20},
		{graph2, 2, 2, -1},
		{graph3, 4, 4, -1},
		{graph4, 5, 3, 59},
	}

	for _, table := range tables {
		v := networkDelayTime(table.i, table.n, table.k)
		t.Logf("v is %d", v)
		if v != table.o {
			t.Errorf("f(g) != %d, but %d", table.o, v)
		}
	}
}
