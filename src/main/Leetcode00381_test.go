package main

import (
	"testing"
)

func TestInsertionDeleteGetRandom(t *testing.T) {
	tables := []struct {
		c []string
		i []int
	}{
		{
			[]string{"insert", "insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "insert", "remove"},
			[]int{9, 9, 1, 1, 2, 1, 2, 1, 1, 9, 1},
		},
		{
			[]string{"insert", "insert", "remove", "getRandom"},
			[]int{1, 1, 1, 0},
		},
	}
	// Shrink array
	for _, table := range tables {
		collection := Constructor()
		for i := range table.c {
			switch table.c[i] {
			case "insert":
				collection.Insert(table.i[i])
			case "remove":
				collection.Remove(table.i[i])
			case "getRandom":
				collection.GetRandom()
			}
		}
	}
}
