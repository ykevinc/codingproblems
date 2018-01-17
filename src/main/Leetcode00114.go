package main

// TreeNode .
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flattenSub(root)
}

func flattenSub(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := flattenSub(root.Left)
	r := flattenSub(root.Right)
	if l != nil {
		l.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if r != nil {
		return r
	} else if l != nil {
		return l
	} else {
		return root
	}
}
