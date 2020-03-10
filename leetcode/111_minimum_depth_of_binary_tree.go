package leetcode

/*
111. Minimum Depth of Binary Tree
Easy

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its minimum depth = 2.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := minDepth(root.Left)
    right := minDepth(root.Right)

    if left == 0 || right == 0 {
        if left > right {
            return left + 1
        }
        return right + 1
    }

    return min(left, right) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
