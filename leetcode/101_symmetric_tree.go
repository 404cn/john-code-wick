package leetcode

/*
101. Symmetric Tree
Easy

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3



But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3



Note:
Bonus points if you could solve it both recursively and iteratively.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return foo(root, root)
}

func foo(rA, rB *TreeNode) bool {
	if rA == nil && rB == nil {
		return true
	}
	if rA == nil || rB == nil {
		return false
	}
	return rA.Val == rB.Val && foo(rA.Left, rB.Right) && foo(rA.Right, rB.Left)
}
