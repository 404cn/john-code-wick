package leetcode

/**
21. Merge Two Sorted Lists
Easy

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	list := &ListNode{}

	if l1.Val < l2.Val {
		list.Val = l1.Val
		list.Next = mergeTwoLists(l1.Next, l2)
	} else {
		list.Val = l2.Val
		list.Next = mergeTwoLists(l1, l2.Next)
	}

	return list
}
