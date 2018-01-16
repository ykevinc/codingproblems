package main

import (
	"reflect"
	"testing"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func TestMerge(t *testing.T) {
	tables := []struct {
		i [][]int
		o [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}
	for _, table := range tables {
		v := mergeInts(table.i)
		if !reflect.DeepEqual(v, table.o) {
			t.Errorf("f(%v) != %v, but %v", table.i, table.o, v)
		}
	}
}
