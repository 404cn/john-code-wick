package leetcode

/*
404. Sum of Left Leaves
Easy

Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	help(root, &sum)
	return sum
}

func help(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	l := root.Left
	if l != nil && l.Left == nil && l.Right == nil {
		*sum += l.Val
	}

	help(root.Left, sum)
	help(root.Right, sum)
}
