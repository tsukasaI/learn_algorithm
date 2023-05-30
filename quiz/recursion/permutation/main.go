package main

import "fmt"

// input: [1, 2, 3]
// output:
// [1, 2, 3]
// [1, 3, 2]
// [2, 1, 3]
// [2, 3, 1]
// [3, 2, 1]
// [3, 1, 2]

func permutation(input []int) [][]int {
	ans := make([][]int, 0)

	if len(input) <= 1 {
		return [][]int{input}
	}
	for _, perm := range permutation(input[1:]) {
		for i := 0; i < len(input); i++ {
			elm := make([]int, 0)
			elm = append(elm, perm[:i]...)
			elm = append(elm, input[0:1]...)
			elm = append(elm, perm[i:]...)
			ans = append(ans, elm)
		}
	}

	return ans
}

func main() {
	input := []int{1, 2, 3, 4}
	ans := permutation(input)
	fmt.Println(ans)
}
