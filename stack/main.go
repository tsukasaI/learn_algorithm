package main

import "fmt"

type stack struct{ stack []int }

func (s *stack) push(data int) {
	s.stack = append(s.stack, data)
}

func (s *stack) pop() *int {
	if len(s.stack) > 0 {
		endIndex := len(s.stack) - 1
		end := s.stack[endIndex]
		s.stack = append(make([]int, 0), s.stack[:endIndex]...)
		return &end
	}
	return nil
}

func validateFormat(chars string) bool {
	lookUp := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	stackList := make([]string, 0)

	for _, char := range chars {
		for key, value := range lookUp {
			if string(char) == key {
				stackList = append(stackList, lookUp[key])
			}
			if string(char) == value {
				if len(stackList) == 0 {
					return false
				}
				endIndex := len(stackList) - 1
				end := stackList[endIndex]
				stackList = append(make([]string, 0), stackList[:endIndex]...)
				if string(char) != end {
					return false
				}
			}
		}
	}

	return len(stackList) == 0
}

func main() {
	s := stack{}
	fmt.Println(s.stack)
	s.push(1)
	fmt.Println(s.stack)
	s.push(2)
	fmt.Println(s.stack)
	fmt.Println(s.pop())
	fmt.Println(s.stack)

	fmt.Println("======quiz")
	j := "{'key1': 'value1', 'key2': 'value2', 'key3': 'value3'}"
	fmt.Println(validateFormat(j))
	k := "{'key1': 'value1', 'key2': 'value2', 'key3': 'value3'}["
	fmt.Println(validateFormat(k))
}
