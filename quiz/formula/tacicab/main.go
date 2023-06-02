package main

// これわかってない No.56

func taxiCabNumber(maxAnswerNum, matchAnswerNum int) {
	result := make([]int, 0)
	gotAnswerCount := 0
	answer := 1

	for gotAnswerCount < matchAnswerNum {
		matchAnswerCount := 0
		memo := make(map[int][]int)

		for i := 1; i < 100; i++ {
			for j := i + 1; j < 100; j++ {
				matchAnswerCount++
				memo[answer] = append(memo[answer], []int{i, j})
			}
		}
	}
}

func main() {
	ans := taxiCabNumber()
}
