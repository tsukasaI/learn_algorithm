package main

import (
	"fmt"
	"strconv"
	"strings"
)

// [1] => [2] => 2
// [2, 3] => [2, 4] => 24
// [8, 9] => [9, 0] => 90
// [1, 2, 3] => [1, 2, 4] => 124
// [7, 8, 9] => [7, 9, 0] => 790

func main() {
	input := []int{9, 9, 9}
	resultStr, resultSlice := addSlice(input)
	fmt.Println(resultSlice)
	fmt.Println(resultStr)
}

func addSlice(input []int) (string, []int) {
	carry := 0
	inputLength := len(input)
	input[inputLength-1]++
	for i := inputLength - 1; i >= 0; i-- {
		input[i] += carry
		if input[i] == 10 {
			carry = 1
			input[i] = 0
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		input[0] = 0
		input = append([]int{1}, input...)
	}

	strSlice := make([]string, len(input))
	for i, v := range input {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, ""), input
}
