package main

import "fmt"

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

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6}

	//fmt.Println(sortedArr, 6, find(sortedArr, 6))
	fmt.Println(sortedArr, 3, find(sortedArr, 0, len(sortedArr), 3))
	fmt.Println(sortedArr, 5, find(sortedArr, 0, len(sortedArr), 5))
	fmt.Println(sortedArr, 5, find(sortedArr, 0, len(sortedArr), 10))
	fmt.Println(sortedArr, 6, find(sortedArr, 0, len(sortedArr), 6))
}
