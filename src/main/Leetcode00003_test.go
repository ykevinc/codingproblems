package main

import (
	"testing"
)

func Test(t *testing.T) {
	tables := []struct {
		i string
		o int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, table := range tables {
		if lengthOfLongestSubstring(table.i) != table.o {
			t.Errorf("f(%s) != %d\n", table.i, table.o)
		}
	}
}
