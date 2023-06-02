package main

import "testing"

func TestIsPrimeV1(t *testing.T) {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19}
	nonPrimeNumbers := []int{1, 4, 8, 12, 14, 18, 20}

	t.Run("OK", func(t *testing.T) {
		for _, primeNumber := range primeNumbers {
			result := isPrimeV1(primeNumber)
			if !result {
				t.Error("expected Prime number, got non prime number")
			}
		}
	})

	t.Run("OK", func(t *testing.T) {
		for _, nonPrimeNumber := range nonPrimeNumbers {
			result := isPrimeV1(nonPrimeNumber)
			if result {
				t.Error("expected non Prime number, got prime number")
			}
		}
	})
}

func TestIsPrimeV2(t *testing.T) {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19}
	nonPrimeNumbers := []int{1, 4, 8, 12, 14, 18, 20}

	t.Run("OK", func(t *testing.T) {
		for _, primeNumber := range primeNumbers {
			result := isPrimeV2(primeNumber)
			if !result {
				t.Error("expected Prime number, got non prime number")
			}
		}
	})

	t.Run("OK", func(t *testing.T) {
		for _, nonPrimeNumber := range nonPrimeNumbers {
			result := isPrimeV2(nonPrimeNumber)
			if result {
				t.Error("expected non Prime number, got prime number")
			}
		}
	})
}

func TestIsPrimeV3(t *testing.T) {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19}
	nonPrimeNumbers := []int{1, 4, 8, 12, 14, 18, 20}

	t.Run("OK", func(t *testing.T) {
		for _, primeNumber := range primeNumbers {
			result := isPrimeV3(primeNumber)
			if !result {
				t.Error("expected Prime number, got non prime number")
			}
		}
	})

	t.Run("OK", func(t *testing.T) {
		for _, nonPrimeNumber := range nonPrimeNumbers {
			result := isPrimeV3(nonPrimeNumber)
			if result {
				t.Error("expected non Prime number, got prime number")
			}
		}
	})
}

func TestIsPrimeV4(t *testing.T) {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19}
	nonPrimeNumbers := []int{1, 4, 8, 12, 14, 18, 20}

	t.Run("OK", func(t *testing.T) {
		for _, primeNumber := range primeNumbers {
			result := isPrimeV4(primeNumber)
			if !result {
				t.Error("expected Prime number, got non prime number")
			}
		}
	})

	t.Run("OK", func(t *testing.T) {
		for _, nonPrimeNumber := range nonPrimeNumbers {
			result := isPrimeV4(nonPrimeNumber)
			if result {
				t.Error("expected non Prime number, got prime number")
			}
		}
	})
}
