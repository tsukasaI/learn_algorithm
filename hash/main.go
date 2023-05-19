package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
)

type hashTable struct {
	size  int
	table [][][]string
}

const defaultSize = 10

func newHashTable() hashTable {
	return hashTable{defaultSize, make([][][]string, defaultSize)}
}

func (h *hashTable) hash(key string) int {
	bi := big.NewInt(0)
	y := big.NewInt(int64(h.size))

	m := md5.New()
	defer m.Reset()
	m.Write([]byte(key))
	str := hex.EncodeToString(m.Sum(nil))
	bi.SetString(str, 16)
	bi.Mod(bi, y)
	return int(bi.Int64())
}

func (h *hashTable) add(key string, value string) {
	isBreaked := false
	index := h.hash(key)
	for _, data := range h.table[index] {
		if data[0] == key {
			data[1] = value
			isBreaked = true
			break
		}
	}
	if !isBreaked {
		h.table[index] = append(h.table[index], []string{key, value})
	}
}

func (h *hashTable) print() {
	for i, data := range h.table {
		fmt.Println(i, data)
	}
}

func (h *hashTable) get(key string) (string, bool) {
	index := h.hash(key)
	for _, data := range h.table[index] {
		if data[0] == key {
			return data[1], true
		}
	}
	return "", false
}

func main() {
	h := newHashTable()

	h.add("car", "Tesla")
	h.add("car", "Tesla")
	h.add("pc", "Mac")
	h.add("sns", "YouTube")
	h.print()
	v, ok := h.get("pc")
	if ok {
		fmt.Println(v)
	}

	fmt.Println("======quiz=====")
	l := []int{11, 2, 5, 9, 10, 3}
	t := 99
	fmt.Println(getPair(l, t))
	fmt.Println("======quiz2=====")

	fmt.Println(getPairHalf(l))
}

func getPair(numbers []int, target int) (*int, *int) {
	cache := make([]int, 0)
	for _, number := range numbers {
		value := target - number

		for _, cached := range cache {
			if cached == value {
				return &value, &number
			}
		}
		cache = append(cache, number)
	}
	return nil, nil
}

func getPairHalf(numbers []int) (*int, *int) {
	sumNumbers := 0
	for _, v := range numbers {
		sumNumbers += v
	}
	if sumNumbers%2 != 0 {
		return nil, nil
	}
	halfSum := sumNumbers / 2

	cache := make([]int, 0)
	for _, number := range numbers {
		cache = append(cache, number)
		value := halfSum - number
		for _, cached := range cache {
			if value == cached {
				return &value, &number
			}
		}
	}
	return nil, nil
}
