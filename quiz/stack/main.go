package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numAlphabetMapping = map[int]string{
	0: "+",
	1: "@",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
}

func phoneNumberMemoric(input string) []string {
	trimed := strings.Replace(input, "-", "", -1)
	digits := strings.Split(trimed, "")
	phoneNumber := make([]int, len(digits))
	for i, digit := range digits {
		numbered, _ := strconv.Atoi(digit)
		phoneNumber[i] = numbered
	}

	candidate := make([]string, 0)
	tmp := make([]string, len(phoneNumber))

	var _findCandidateAlphabet func(digit int)
	_findCandidateAlphabet = func(digit int) {
		if digit == len(phoneNumber) {
			candidate = append(candidate, strings.Join(tmp, ""))
		} else {
			numAlphabet := numAlphabetMapping[phoneNumber[digit]]
			numAlphabetSlice := strings.Split(numAlphabet, "")
			for _, char := range numAlphabetSlice {
				tmp[digit] = char
				_findCandidateAlphabet(digit + 1)
			}
		}
	}

	_findCandidateAlphabet(0)
	return candidate
}

func phoneNumberMemoricStack(input string) []string {
	trimed := strings.Replace(input, "-", "", -1)
	digits := strings.Split(trimed, "")
	phoneNumber := make([]int, len(digits))
	for i, digit := range digits {
		numbered, _ := strconv.Atoi(digit)
		phoneNumber[i] = numbered
	}

	candidate := make([]string, 0)
	stack := make([]string, 1)

	for len(stack) != 0 {
		alphabets := stack[len(stack)-1]
		stack = append(make([]string, 0), stack[:len(stack)-1]...)

		if len(alphabets) == len(phoneNumber) {
			candidate = append(candidate, alphabets)
		} else {
			numAlphabet := numAlphabetMapping[phoneNumber[len(alphabets)]]
			numAlphabetSlice := strings.Split(numAlphabet, "")

			for _, char := range numAlphabetSlice {
				stack = append(stack, alphabets+char)
			}
		}
	}

	return candidate
}

func main() {
	// input := "568-379-8466"
	input := "23"
	ans := phoneNumberMemoricStack(input)
	fmt.Println(ans)
}
