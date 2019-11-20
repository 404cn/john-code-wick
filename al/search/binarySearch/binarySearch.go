// 二分查找, 折半查找 o(logn)
// 有序元素， 连续存放
// 循环推出条件，mid的取值，low与high的更新
package main

import "fmt"

func main() {
	// 错误示范
	// arr := []int{5, 4, 3, 2, 1}
	arr := []int{1, 2, 3, 4, 5}
	target := 4
	fmt.Println(binarySearch(arr, target))
	fmt.Println(rbs(arr, target, 0, len(arr)-1))
}

// find target's position in numbers
// 简单二分查找，不考虑 重复元素
func binarySearch(numbers []int, target int) int {
	low, high := 0, len(numbers)-1
	for low <= high {
		// 可能造成溢出，改进并用位运算代替除法
		//mid := (low + high) / 2
		mid := low + (high-low)>>1
		switch {
		case numbers[mid] == target:
			return mid
		case numbers[mid] > target:
			high = mid - 1
		case numbers[mid] < target:
			low = mid + 1
		}
	}
	return -1
}

// recursion version
func rbs(numbers []int, target, low, high int) int {
	// 退出条件
	if low > high || len(numbers) == 0 {
		return -1
	}
	mid := low + (high-low)>>1
	if numbers[mid] == target {
		return mid
	} else if numbers[mid] > target {
		high = mid - 1
	} else if numbers[mid] < target {
		low = mid + 1
	}
	return rbs(numbers, target, low, high)
}
