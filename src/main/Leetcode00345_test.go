package main

import (
	"testing"
)

func TestRverseVowels(t *testing.T) {
	tables := []struct {
		i string
		o string
	} {
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}
	for _, table := range tables {
		v := reverseVowels(table.i)
		if v != table.o {
			t.Errorf("f(%s) != %s, but %s", table.i, table.o, v)
		}
	}
}
