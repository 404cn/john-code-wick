package leetcode

/*
993. Cousins in Binary Tree
Easy

In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.



Example 1:

Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Example 2:

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true

Example 3:

Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false



Note:

    The number of nodes in the tree will be between 2 and 100.
    Each node has a unique integer value from 1 to 100.
*/

func isCousins(root *TreeNode, x int, y int) bool {
	depth := make(map[int]int)
	parent := make(map[int]*TreeNode)
	dfs(root, nil, depth, parent)
	return depth[x] == depth[y] && parent[x] != parent[y]
}

func dfs(root, par *TreeNode, depth map[int]int, parent map[int]*TreeNode) {
	if root != nil {
		if par == nil {
			depth[root.Val] = 0
		} else {
			depth[root.Val] = 1 + depth[par.Val]
		}
		parent[root.Val] = par
		dfs(root.Left, root, depth, parent)
		dfs(root.Right, root, depth, parent)
	}
}
