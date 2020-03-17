package leetcode

// sort
/*
349. Intersection of Two Arrays
Easy

635

1048

Add to List

Share
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
*/

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m := make(map[int]int)
	var a []int

	for _, i := range nums1 {
		m[i] = 0
	}

	for _, i := range nums2 {
		if _, ok := m[i]; ok {
			m[i]++
		}
	}

	for k, v := range m {
		if v != 0 {
			a = append(a, k)
		}
	}
	return a
}

// 需要考虑的:
// 重复项,负值,单值列表,0与空列表参数
