package main

import (
	"math/rand"
)

// RandomizedCollection .
type RandomizedCollection struct {
	m map[int]map[int]bool
	a []int
}

// Constructor .
func Constructor() RandomizedCollection {
	return RandomizedCollection{make(map[int]map[int]bool, 20), make([]int, 0, 20)}
}

// Insert .
func (t *RandomizedCollection) Insert(val int) bool {
	v := true
	if len(t.m[val]) > 0 {
		v = false
	}
	w := len(t.a)
	if t.m[val] == nil {
		t.m[val] = make(map[int]bool)
	}
	t.m[val][w] = true
	t.a = append(t.a, val)
	return v
}

func (t *RandomizedCollection) pollKey(val int) (int, bool) {
	for k := range t.m[val] {
		delete(t.m[val], k)
		return k, true
	}
	return 0, false
}

// Remove .
func (t *RandomizedCollection) Remove(val int) bool {
	p, ok := t.pollKey(val)
	if !ok {
		return false
	}
	// Swap
	t.a[p] = t.a[len(t.a)-1]
	t.a = t.a[:len(t.a)-1]
	// Remove from map
	if p != len(t.a) {
		delete(t.m[t.a[p]], len(t.a))
		t.m[t.a[p]][p] = true
	}
	return true
}

// GetRandom .
func (t *RandomizedCollection) GetRandom() int {
	return t.a[rand.Intn(len(t.a))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
