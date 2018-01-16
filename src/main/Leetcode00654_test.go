package main

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tables := []struct {
		i []int
		o []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{3, 1, 1}, []int{1, 2}},
		{[]int{1, 2, 3, 3}, []int{3, 4}},
	}
	for _, table := range tables {
		v := findErrorNums(table.i)
		if !reflect.DeepEqual(v, table.o) {
			t.Errorf("f(%v) != %v, but %v", table.i, table.o, v)
		}
	}
}
