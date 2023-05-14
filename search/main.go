package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 8, 9, 11, 15, 18, 22, 24}

	fmt.Println("------linearSearch------")
	fmt.Println(linearSearch(nums, 4))
	fmt.Println(linearSearch(nums, 100))
	fmt.Println("------binarySearch------")
	fmt.Println(binarySearch(nums, 4))
	fmt.Println(binarySearch(nums, 100))
	fmt.Println("------binarySearchRecursive------")
	fmt.Println(binarySearchRecursive(nums, 4))
	fmt.Println(binarySearchRecursive(nums, 100))
}

func linearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}

func binarySearch(numbers []int, target int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		center := (left + right) / 2
		if numbers[center] == target {
			return center
		} else if numbers[center] < target {
			left = center + 1
		} else {
			right = center - 1
		}
	}
	return -1
}

func binarySearchRecursive(numbers []int, target int) int {
	var _binarySearch func(numbers []int, target, left, right int) int
	_binarySearch = func(numbers []int, target, left, right int) int {
		if left > right {
			return -1
		}
		center := (left + right) / 2
		if numbers[center] == target {
			return center
		} else if numbers[center] < target {
			return _binarySearch(numbers, target, center+1, right)
		} else {
			return _binarySearch(numbers, target, left, center-1)
		}
	}
	return _binarySearch(numbers, target, 0, len(numbers)-1)
}
