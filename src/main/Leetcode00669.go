package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}

	val := root.Val
	if val >= L && val <= R {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
	tmp := root
	if val < L {
		tmp = root.Right
	} else if val > R {
		tmp = root.Left
	}
	return trimBST(tmp, L, R)
}
