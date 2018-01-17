package main

import (
	"testing"
)

type CartOffer struct {
	p []int
	s [][]int
	n []int
}

func TestShoppingOffers(t *testing.T) {
	tables := []struct {
		i CartOffer
		o int
	}{
		{CartOffer{[]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}}, 11},
		{CartOffer{[]int{4,3,2,9,8,8}, [][]int{{1,5,5,1,4,0,18},{3,3,6,6,4,2,32}}, []int{6,5,5,6,4,1}}, 91},
	}
	for _, table := range tables {
		v := shoppingOffers(table.i.p, table.i.s, table.i.n)
		if v != table.o {
			t.Errorf("f(%v) != %d, but %d", table.i, table.o, v)
		}
	}

}
