package main

import (
	"fmt"
	"strconv"
)

func showSnakeNumbers(lineWidth, length int) {
	lines := make([][]string, lineWidth)

	position := lineWidth / 2
	isToUp := true
	for i := 0; i < length; i++ {
		for j := 0; j < lineWidth; j++ {
			if j == position {
				lines[j] = append(lines[j], strconv.Itoa(i))
			} else {
				lines[j] = append(lines[j], "x")
			}

		}
		if position == 0 {
			isToUp = false
		} else if position == lineWidth-1 {
			isToUp = true
		}

		if isToUp {
			position--
		} else {
			position++
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	showSnakeNumbers(5, 20)
}
