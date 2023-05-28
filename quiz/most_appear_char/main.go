// Input: "This is a pen. This is an apple/ Applepen."
// Output: ("p", 6)

package main

import (
	"fmt"
	"strings"
)

func getMostAppearedChar(input string) [2]any {
	hm := make(map[rune]int)
	allLowerStr := strings.ToLower(input)
	whiteSpaceTrimedStr := strings.Replace(allLowerStr, " ", "", -1)
	for _, char := range whiteSpaceTrimedStr {
		hm[char]++
	}

	fmt.Println(hm)

	var resultKey rune
	maxNum := 0
	for key, value := range hm {
		if maxNum < value {
			resultKey = key
			maxNum = value
		}
	}
	return [2]any{string(resultKey), maxNum}
}

func main() {
	input := "This is a pen. This is an apple/ Applepen."
	ans := getMostAppearedChar(input)
	fmt.Println(ans)
}
