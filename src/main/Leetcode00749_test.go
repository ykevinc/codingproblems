package main

import (
	"testing"
)

func TestContainVirus(t *testing.T) {
	grid1 := [][]int{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	grid2 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}
	grid3 := [][]int{
		{1, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0}}

	tables := []struct {
		grid [][]int
		o    int
	}{
		{grid1, 10},
		{grid2, 4},
		{grid3, 13},
	}
	for i, table := range tables {
		v := containVirus(table.grid)
		if v != table.o {
			t.Errorf("f(%v) != %d, but %d", i, table.o, v)
		}
	}
}
