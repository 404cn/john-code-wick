// Package main provides ...
package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for k, v := range numbers {
		if idx, ok := m[target-v]; ok && idx < k {
			return []int{idx + 1, k + 1}
		}
		m[v] = k
	}
	return nil
}
