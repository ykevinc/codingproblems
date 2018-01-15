package main

import (
	"math/bits"
)

func countBits(num int) []int {
	num64 := uint64(num)
	l := bits.Len64(num64);
	a := make([]int, 0, l)
	for i := uint64(0);i <= num64;i++ {
		a = append(a, bits.OnesCount64(i))
	}
	return a
}