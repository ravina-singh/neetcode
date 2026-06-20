/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func height(root *TreeNode, balanced *bool) int{
	if root == nil {
		return 0
	}
	
	lh := height(root.Left, balanced)
	rh := height(root.Right, balanced)

	*balanced = *balanced && abs(lh-rh) <= 1
	return 1 + max(lh, rh)
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	balanced := true
	height(root, &balanced)
	return balanced
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
