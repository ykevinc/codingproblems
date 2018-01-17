package main

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	s := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	flatten(s)
	if s.Val != 1 {
		t.Errorf("s.val is incorrect %d", s.Val)
	}
	if s.Left != nil {
		t.Errorf("s.Left is incorrect %d", s.Left.Val)
	}
	if s.Right.Val != 2 {
		t.Errorf("s.Right.val is incorrect %d", s.Right.Val)
	}
	if s.Right.Left != nil {
		t.Errorf("s.Right.Left is incorrect %d", s.Right.Left.Val)
	}
	if s.Right.Right.Val != 3 {
		t.Errorf("s.Right.Right.Val is incorrect %d", s.Right.Val)
	}
}
