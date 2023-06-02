package main

import "fmt"

// input: 50 => output[2, 3, 4, 7, 11, 13, 17, 23, 29, 31, 37, 41, 43, 47]

func generatePrimeNumbersV1(input int) []int {
	primes := make([]int, 0)
	mod0 := false
	for i := 2; i <= input; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				mod0 = true
			}
		}

		if mod0 {
			mod0 = false
			continue
		}
		primes = append(primes, i)
	}
	return primes
}

func generatePrimeNumbersV2(input int) []int {
	primes := make([]int, 0)
	cache := make(map[int]bool, 0)

	for i := 2; i <= input; i++ {
		_, ok := cache[i]
		if ok {
			continue
		}
		cache[i] = false
		primes = append(primes, i)
		for j := i; j <= input; j += i {
			cache[j] = false
		}
	}
	return primes
}

func main() {
	input := 50
	ans := generatePrimeNumbersV1(input)
	fmt.Println(ans)
	ans2 := generatePrimeNumbersV2(input)
	fmt.Println(ans2)
}
