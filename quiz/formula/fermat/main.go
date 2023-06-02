package main

import (
	"fmt"
	"math"
)

func fermatLastTheoremV1(maxNum, squireNum int) [][]int {
	result := make([][]int, 0)

	if squireNum < 2 {
		return result
	}

	maxZ := int(math.Pow(float64((maxNum-1)*(maxNum-1)+maxNum*maxNum), 1.0/2))

	for i := 1; i <= maxNum; i++ {
		for j := i + 1; j <= maxNum; j++ {
			for k := j + 1; k <= maxZ; k++ {
				if math.Pow(float64(i), float64(squireNum))+math.Pow(float64(j), float64(squireNum)) == math.Pow(float64(k), float64(squireNum)) {
					result = append(result, []int{i, j, k})
				}
			}
		}
	}

	return result
}

func fermatLastTheoremV2(maxNum, squireNum int) [][]int {
	result := make([][]int, 0)

	if squireNum < 2 {
		return result
	}

	for i := 1; i <= maxNum; i++ {
		for j := i + 1; j <= maxNum; j++ {
			powSum := math.Pow(float64(i), float64(squireNum)) + math.Pow(float64(j), float64(squireNum))
			k := int(math.Pow(powSum, 1.0/float64(squireNum)))
			kPow := math.Pow(float64(k), float64(squireNum))
			if kPow == powSum {
				result = append(result, []int{i, j, k})
			}
		}
	}

	return result
}

func main() {
	fmt.Println("v1", fermatLastTheoremV1(20, 3))
	fmt.Println("v1", fermatLastTheoremV2(20, 2))
}
