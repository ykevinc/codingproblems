package main

func findTarget(root *TreeNode, k int) bool {
   return findTargetSub(root, k, make(map[int]bool, 10)) 
}

func findTargetSub(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}
	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true
	if findTargetSub(root.Left, k, m) ||  findTargetSub(root.Right, k, m)  {
		return true
	}
	return false
}