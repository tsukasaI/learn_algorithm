package main

import "fmt"

// 1. Check palindrome
// aba -> true
// abc -> false
// racecar -> true

// 2. Find palindrome
// abcracecarbda -> cec, aceca, racecar

func checkPalidrome(input string) bool {
	if len(input) <= 1 {
		return true
	}
	runedInput := []rune(input)

	if runedInput[0] == runedInput[len(runedInput)-1] {
		str := string(runedInput[1 : len(runedInput)-1])
		return checkPalidrome(str)
	} else {
		return false
	}
}

func findPalindrome(str string, left, right int) []string {
	runedStr := []rune(str)
	result := make([]string, 0)
	for left >= 0 && right <= len(runedStr)-1 {
		if runedStr[left] != runedStr[right] {
			break
		}
		result = append(result, string(runedStr[left:right+1]))
		left--
		right++
	}
	return result
}

func findAllPalindrome(str string) []string {
	result := make([]string, 0)
	lenStrings := len(str)
	if lenStrings == 0 {
		return result
	}
	if lenStrings == 1 {
		result = append(result, str)
		return result
	}

	for i := 0; i < lenStrings-1; i++ {
		result = append(result, findPalindrome(str, i-1, i+1)...)
		result = append(result, findPalindrome(str, i-1, i)...)
	}
	return result
}

func main() {
	input := "aaracecar"
	result1 := checkPalidrome(input)
	fmt.Println(result1)

	result2 := findAllPalindrome(input)
	fmt.Println(result2)

}
