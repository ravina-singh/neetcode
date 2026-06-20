/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func height(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left, res)
	rh := height(root.Right, res)
	*res = max(*res, lh+rh)
	return 1 + max(lh, rh)
}

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}
	res := 0
	height(root, &res)
	return res
}
