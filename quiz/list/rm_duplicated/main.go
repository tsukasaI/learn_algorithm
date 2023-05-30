package main

import (
	"fmt"
	"sort"
)

// input [1, 3, 3, 4, 4, 5, 7, 7, 7, 10, 12, 15]
// output [1, 3, 4, 5, 7, 10, 12, 15]

func rmDuplicated(input []int) []int {
	hm := make(map[int]bool)
	for _, v := range input {
		hm[v] = true
	}
	ans := make([]int, 0)
	for i := range hm {
		ans = append(ans, i)
	}
	sort.Ints(ans)
	return ans
}

func main() {
	input := []int{1, 3, 3, 4, 4, 5, 7, 7, 7, 10, 12, 15}
	result := rmDuplicated(input)
	fmt.Println(result)
}
