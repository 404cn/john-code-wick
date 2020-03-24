package leetcode

/*
501. Find Mode in Binary Search Tree
Easy

749

288

Add to List

Share
Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.


For example:
Given BST [1,null,2,2],

   1
    \
     2
    /
   2


return [2].

Note: If a tree has more than one mode, you can return them in any order.

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	m := make(map[int]int)
	a := make([]int, 0)

	help(root, m)

	max := 0

	for _, v := range m {
		if v >= max {
			max = v
		}
	}

	for k, _ := range m {
		if m[k] == max {
			a = append(a, k)
		}
	}

	return a
}

func help(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	m[root.Val]++

	help(root.Left, m)
	help(root.Right, m)
}

/*
without extra space

func findMode(root *TreeNode) []int {
    if root == nil { return nil }
    prev := math.MinInt32
    count := 0
    max := 1
    res := []int{}
    helper(root, &prev, &count, &max, &res)
    return res
}

func helper(root *TreeNode, prev, count, max *int, res *[]int) {
    if root == nil { return }
    helper(root.Left, prev, count, max, res)
    if root.Val == *prev { *count++ } else { *count = 1 }
    if *max < *count {
        *res = []int{root.Val}
        *max = *count
    } else if *max == *count {
        *res = append(*res, root.Val)
    }
    *prev = root.Val
    helper(root.Right, prev, count, max, res)
}

*/
