package main

import (
	"fmt"
)

// 1. Max subarray sum
// input: [1, -2, 3, 6, -1, 2, 4, -5, 2]
// output: 14 (3, 6, -1, 2, 4)

// 2. Max circular subarray sum
// input: [1, -2, 3, 6, -1, 2, 4, -5, 2]
// output: 15 (2, 1, 2, 3, 6, -1, 2, 4)

func getMaxSequenceSum(numbers []int) int {
	resultSequence, sumSequence := 0, 0
	for _, num := range numbers {
		tempSum := sumSequence + num
		if num < tempSum {
			sumSequence = tempSum
		} else {
			sumSequence = num
		}

		if resultSequence < sumSequence {
			resultSequence = sumSequence
		}
	}
	return resultSequence
}

func findMaxCircularSequence(numbers []int) int {
	maxSequenceSum := getMaxSequenceSum(numbers)

	invertNumbers := make([]int, 0)
	allSum := 0
	for _, num := range numbers {
		allSum += num
		invertNumbers = append(invertNumbers, -num)
	}

	maxWrapSequence := allSum - (-getMaxSequenceSum(invertNumbers))

	if maxSequenceSum < maxWrapSequence {
		return maxWrapSequence
	} else {
		return maxSequenceSum
	}
}

func main() {
	ans := getMaxSequenceSum([]int{1, -2, 3, 6, -1, 2, 4, -5, 2})
	fmt.Println(ans)

	ans2 := findMaxCircularSequence([]int{1, -2, 3, 6, -1, 2, 4, -5, 2})
	fmt.Println(ans2)

}
