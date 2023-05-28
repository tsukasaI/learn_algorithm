package main

import "fmt"

func symmetric(input [][2]int) [][2]int {
	ans := make([][2]int, 0)
	if len(input) < 2 {
		return [][2]int{}
	}

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i][0] == input[j][1] && input[i][1] == input[j][0] {
				ans = append(ans, input[j])
				continue
			}
		}
	}
	return ans
}

func findPair(input [][2]int) [][2]int {
	hm := make(map[int]int)
	ans := make([][2]int, 0)
	for _, tuple := range input {
		if value, ok := hm[tuple[0]]; ok && value == tuple[1] {
			ans = append(ans, tuple)
		} else {
			hm[tuple[1]] = tuple[0]
		}
	}
	return ans
}

func main() {
	input := [][2]int{
		[2]int{1, 2},
		[2]int{3, 5},
		[2]int{4, 7},
		[2]int{5, 3},
		[2]int{7, 4},
	}

	result := findPair(input)
	fmt.Println(result)
}
