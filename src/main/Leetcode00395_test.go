package main

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	table := []struct {
		i string
		k int
		o int
	}{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
	}
	for _, table := range table {
		v := longestSubstring(table.i, table.k)
		if v != table.o {
			t.Errorf("f(%s) != %d, but %d", table.i, table.o, v)
		}
	}
}
