package leetcode

import "strconv"

/*

257. Binary Tree Paths
Easy

Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	paths := make([]string, 0)
	help(root, strconv.Itoa(root.Val), &paths)
	return paths
}

func help(root *TreeNode, path string, paths *[]string) {
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
	}

	if root.Left != nil {
		help(root.Left, path+"->"+strconv.Itoa(root.Left.Val), paths)
	}

	if root.Right != nil {
		help(root.Right, path+"->"+strconv.Itoa(root.Right.Val), paths)
	}
}
