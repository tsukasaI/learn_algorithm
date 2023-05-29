package main

import "fmt"

// input: X: [1, 2, 3, 4, 4, 5, 5, 8, 10], Y: [4, 5, 5, 5, 6, 7, 8, 8, 10]
// output X: [1, 2, 3, 4, 4, 10], Y: [5, 5, 5, 6, 7, 8, 8, 10]

func main() {
	X := []int{1, 2, 3, 4, 4, 5, 5, 8, 10}
	Y := []int{4, 5, 5, 5, 6, 7, 8, 8, 10}

	resultX, resultY := minCountRemove(X, Y)
	fmt.Println(resultX)
	fmt.Println(resultY)
}

func minCountRemove(X, Y []int) ([]int, []int) {
	countX, countY := make(map[int]int), make(map[int]int)

	for _, value := range X {
		countX[value]++
	}
	for _, value := range Y {
		countY[value]++
	}

	for keyX, valueX := range countX {
		valueY, ok := countY[keyX]
		if ok {
			if valueX < valueY {
				newX := make([]int, 0)
				for _, v := range X {
					if v != keyX {
						newX = append(newX, v)
					}
					X = newX
				}
			} else if valueX > valueY {
				newY := make([]int, 0)
				for _, v := range Y {
					if v != keyX {
						newY = append(newY, v)
					}
					Y = newY
				}
			}

		}
	}
	return X, Y
}
