package main

import "fmt"

// input: [0, 1, 3, 4, 2, 5, 1, 6, 9, 8]
// output: [0, 4, 2, 6, 8, 1, 3, 5, 1, 9]

func orderEvenFirstOddLastV1(input []int) {
	evenList, oddList := make([]int, 0), make([]int, 0)
	for _, v := range input {
		if v%2 == 0 {
			evenList = append(evenList, v)
		} else {
			oddList = append(oddList, v)
		}
	}
	input = append(evenList, oddList...)
}
func orderEvenFirstOddLastV2(input []int) {
	i, j := 0, len(input)-1
	for i < j {
		if input[i]%2 == 0 {
			i++
		} else {
			for input[i]%2 != 0 {
				input[i], input[j] = input[j], input[i]
				j--
			}
			i++
		}
	}
}

func main() {
	input := []int{0, 1, 3, 4, 2, 5, 1, 6, 9, 8}
	orderEvenFirstOddLastV2(input)
	fmt.Println(input)
}
