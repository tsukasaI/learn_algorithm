package main

import (
	"fmt"
)

var (
	runedA = []rune("a")[0]
	runedZ = []rune("z")[0]
)

func generateKey(message, keyword string) string {
	key := keyword
	remainLength := len(message) - len(keyword)
	for i := 0; i < remainLength; i++ {
		key += string(key[i])
	}
	return key
}

func encrypt(message, key string) string {
	result := ""
	runedMessage, runedKey := []rune(message), []rune(key)
	for i, runedChar := range runedMessage {
		if runedChar < runedA || runedChar > runedZ {
			result += string(runedChar)
			continue
		}

		targetRuneIndex := ((runedChar - runedA) + runedKey[i]) % (runedZ - runedA + 1)
		fmt.Println(targetRuneIndex)
		result += string(runedA + targetRuneIndex)
	}
	return result
}

func main() {
	t := "attack silicon valley"
	key := generateKey(t, "hello")
	e := encrypt(t, key)
	fmt.Println(e)
}
