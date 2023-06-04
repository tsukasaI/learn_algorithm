package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generatePascalsTriangle(head, depth int) [][]int {
	data := make([][]int, depth)
	for i := 0; i < depth; i++ {
		data[i] = make([]int, i+1)
	}
	data[0][0] = head
	data[1][0], data[1][1] = head, head

	for line := 2; line < depth; line++ {
		data[line][0] = head
		for i := 1; i < line; i++ {
			data[line][i] = data[line-1][i-1] + data[line-1][i]
		}
		data[line][line] = head
	}
	return data
}

func printPascal(data [][]int) {
	for _, line := range data {
		strLine := make([]string, len(line))
		for j := range line {
			strLine[j] = strconv.Itoa(line[j])
		}

		fmt.Println(strings.Join(strLine, ""))

	}
}

func main() {
	result := generatePascalsTriangle(1, 5)
	printPascal(result)
}
