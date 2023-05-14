package main

import (
	"fmt"
	"math/rand"
)

const length = 100

func main() {
	originalNumbers := getRandomSlice()

	fmt.Println("--------original--------")
	fmt.Println(originalNumbers)

	numbers := make([]int, len(originalNumbers))

	// fmt.Println("--------bodoSort--------")
	// copy(numbers, originalNumbers)
	// bogoSort(numbers)
	// fmt.Println(numbers)

	fmt.Println("--------bubbleSort--------")
	copy(numbers, originalNumbers)
	bubbleSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------cocktailSort--------")
	copy(numbers, originalNumbers)
	cocktailSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------combSort--------")
	copy(numbers, originalNumbers)
	combSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------selectionSort--------")
	copy(numbers, originalNumbers)
	selectionSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------gnomeSort--------")
	copy(numbers, originalNumbers)
	gnomeSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------insertionSort--------")
	copy(numbers, originalNumbers)
	insertionSort(numbers)
	fmt.Println(numbers)

	// fmt.Println("--------bucketSort--------")
	// copy(numbers, originalNumbers)
	// bucketSort(numbers)
	// fmt.Println(numbers)

	fmt.Println("--------quickSort--------")
	copy(numbers, originalNumbers)
	quickSort(numbers)
	fmt.Println(numbers)

	fmt.Println("--------mergeSort--------")
	copy(numbers, originalNumbers)
	mergeSortResult := mergeSort(numbers)
	fmt.Println(mergeSortResult)
}

func getRandomSlice() []int {
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i
	}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i] //スライスの要素を入れ替える
	})
	return slice
}

func bogoSort(numbers []int) {
	for !isInOrders(numbers) {
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i] //スライスの要素を入れ替える
		})
	}
}

func isInOrders(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func cocktailSort(numbers []int) {
	lenNumbers := len(numbers)
	swapped := true
	start := 0
	end := lenNumbers - 1
	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
			if !swapped {
				break
			}
			swapped = false
			end--
		}

		for i := end - 1; i > start-1; i-- {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}
		start++
	}
}

func combSort(numbers []int) {
	lenNumbers := len(numbers)
	gap := lenNumbers
	swapped := true

	for gap != 1 || swapped {
		gap = int(float64(gap) / 1.3)
		if gap < 1 {
			gap = 1
		}
		swapped = false
		for i := 0; i < lenNumbers-gap; i++ {
			if numbers[i] > numbers[i+gap] {
				numbers[i], numbers[i+gap] = numbers[i+gap], numbers[i]
				swapped = true
			}
		}
	}
}

func selectionSort(numbers []int) {
	lenNumbers := len(numbers)
	for start := 0; start < lenNumbers; start++ {
		tmpIndex := start
		for i := start; i < lenNumbers; i++ {
			if numbers[tmpIndex] > numbers[i] {
				tmpIndex = i
			}
		}
		if tmpIndex != start {
			numbers[tmpIndex], numbers[start] = numbers[start], numbers[tmpIndex]
		}
	}
}

func gnomeSort(numbers []int) {
	lenNumbers := len(numbers)
	index := 0
	for index < lenNumbers {
		if index == 0 {
			index++
		}
		if numbers[index] >= numbers[index-1] {
			index++
		} else {
			numbers[index], numbers[index-1] = numbers[index-1], numbers[index]
			index--
		}
	}

}

func insertionSort(numbers []int) {
	lenNumbers := len(numbers)
	for i := 1; i < lenNumbers; i++ {
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}
}

func bucketSort(numbers []int) {
	max := 0
	for _, value := range numbers {
		if max <= value {
			max = value
		}
	}
	lenNumbers := len(numbers)
	size := max / lenNumbers
	buckets := make([][]int, size)

	for _, value := range numbers {
		i := value / size
		if i != size {
			buckets[i] = append(buckets[i], value)
		} else {
			buckets[size-1] = append(buckets[size-1], value)
		}
	}

	for i := 0; i < size; i++ {
		insertionSort(buckets[i])
	}

	result := make([]int, 0)
	for _, value := range buckets {
		result = append(result, value...)
	}
}

func quickSort(numbers []int) {
	var _quickSort func(numbers []int, low, high int)
	_quickSort = func(numbers []int, low, high int) {
		if low < high {
			partitionIndex := partition(numbers, low, high)
			_quickSort(numbers, low, partitionIndex-1)
			_quickSort(numbers, partitionIndex+1, high)
		}
	}
	_quickSort(numbers, 0, len(numbers)-1)
}

func partition(numbers []int, low, high int) int {
	i := low - 1
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func mergeSort(numbers []int) []int {
	result := make([]int, 0)
	if len(numbers) <= 1 {
		return numbers
	}
	center := len(numbers) / 2
	left := numbers[:center]
	right := numbers[center:]

	r := mergeSort(left)
	l := mergeSort(right)
	i, j := 0, 0

	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}

	for i < len(l) {
		result = append(result, l[i])
		i++
	}

	for j < len(r) {
		result = append(result, r[j])
		j++
	}
	return result
}
