package main

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	tables := []struct {
		w []string
		b [][]byte
		o []string
	}{
		{
			[]string{"oath", "pea", "eat", "rain"},
			[][]byte{{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'}},
			[]string{"eat", "oath"},
		},
	}

	for _, table := range tables {
		v := findWords(table.b, table.w)

		if !reflect.DeepEqual(v, table.o) {
			t.Errorf("f(%v) != %v, but %v", table.w, table.o, v)
		}
	}
}
