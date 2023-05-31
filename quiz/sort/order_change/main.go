package main

import (
	"fmt"
	"strings"
)

// input: ["h", "y", "n", "p", "t", "o"], [3, 1, 5, 0, 2, 4]
// output: python
func orderChangeByIndexes(chars []string, indexes []int) string {
	if len(chars) != len(indexes) {
		return ""
	}
	resultChars := make([]string, len(chars))
	for i, index := range indexes {
		resultChars[index] = chars[i]
		fmt.Println(resultChars)
	}
	return strings.Join(resultChars, "")
}

// in place
func orderChangeByIndexesV2(chars []string, indexes []int) string {
	if len(chars) != len(indexes) {
		return ""
	}

	i := 0

	for i < len(chars) {
		if i != indexes[i] {
			chars[i], chars[indexes[i]] = chars[indexes[i]], chars[i]
			indexes[i], indexes[indexes[i]] = indexes[indexes[i]], indexes[i]
		}

		if i == indexes[i] {
			i++
		}
	}
	return strings.Join(chars, "")
}

func main() {
	strInput := []string{"h", "y", "n", "p", "t", "o"}
	indexInput := []int{3, 1, 5, 0, 2, 4}
	ans := orderChangeByIndexes(strInput, indexInput)
	fmt.Println(ans)
	v2 := orderChangeByIndexesV2(strInput, indexInput)
	fmt.Println(v2)
}
