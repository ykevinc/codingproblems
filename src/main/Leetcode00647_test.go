package main

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
    tables := []struct {
		s string
		o int
	} {
		{"abc", 3},
		{"aaa", 6},
	}
	for _, table := range tables {
		v := countSubstrings(table.s)
		if v != table.o {
			t.Errorf("f(%s) != %d, but %d", table.s, table.o, v)
		}
	}
}