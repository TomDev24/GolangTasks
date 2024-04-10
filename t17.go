package main

import "fmt"

// 1. Recursive method
func find(arr []int, l, h, el int) int {
	if l == h {
		return -1
	}
	mid := (l + h) / 2

	if el < arr[mid] {
		return find(arr, l, mid, el)
	} else if el > arr[mid] {
		return find(arr, mid + 1, h, el)
	} else {
		return mid
	}
}

// 2. Iterative method
func binarySearch(data []int, target int) (int, bool) {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2

		if target == data[mid] {
			return mid, true
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6}

	//Recursive
	fmt.Println(sortedArr, 3, find(sortedArr, 0, len(sortedArr), 3))
	fmt.Println(sortedArr, 5, find(sortedArr, 0, len(sortedArr), 5))
	fmt.Println(sortedArr, 10, find(sortedArr, 0, len(sortedArr), 10))
	fmt.Println(sortedArr, 6, find(sortedArr, 0, len(sortedArr), 6))

	//Iterative
	i, ok := binarySearch(sortedArr, 5)
	fmt.Println(sortedArr, 5, i, ok)
}
