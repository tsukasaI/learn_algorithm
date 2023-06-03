package main

import (
	"fmt"
	"unicode"

	"golang.org/x/exp/slices"
)

func caesarCipher(text string, shift int) string {
	result, runedAlphabet := "", make([]rune, 0)

	for _, char := range text {
		if unicode.IsUpper(char) {
			for cur := []rune("A"); ; cur[0]++ {
				runedAlphabet = append(runedAlphabet, cur[0])
				if string(cur[0]) == "Z" {
					break
				}
			}
		} else if unicode.IsLower(char) {
			for cur := []rune("a"); ; cur[0]++ {
				runedAlphabet = append(runedAlphabet, cur[0])
				if string(cur[0]) == "z" {
					break
				}
			}
		} else {
			result += string(char)
			continue
		}

		i := (slices.Index[rune](runedAlphabet, []rune(string(char))[0]) + shift) % len(runedAlphabet)
		result += string(runedAlphabet[i])
	}
	return result
}

func main() {
	ans := caesarCipher("ATTACKED by a x y z", 1)
	fmt.Println(ans)
	fmt.Println([]rune("a"))
	fmt.Println([]rune("A"))
	fmt.Println([]rune("c"))
}
