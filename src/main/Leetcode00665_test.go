package main

import (
	"testing"
)

func TestCheckPossibility(t *testing.T) {
	tables := []struct{
		i []int
		o bool
	} {
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{2,3,3,2,4}, true},
	}
	for _, table := range tables {
		v := checkPossibility(table.i);
		if v != table.o {
			t.Errorf("f(%v) != %t, but %t", table.i, table.o, v) 
		}
	}
}