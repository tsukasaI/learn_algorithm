package main

import (
	"fmt"
)

type miniHeap struct {
	heap []int
	size int
}

func (mh *miniHeap) parentIndex(index int) int {
	return index / 2
}

func (mh *miniHeap) leftChildIndex(index int) int {
	return index * 2
}

func (mh *miniHeap) rightChildIndex(index int) int {
	return (index * 2) + 1
}

func (mh *miniHeap) swap(index1, index2 int) {
	mh.heap[index1], mh.heap[index2] = mh.heap[index2], mh.heap[index1]
}

func (mh *miniHeap) minChildIndex(index int) int {
	if mh.rightChildIndex(index) > mh.size {
		return mh.leftChildIndex(index)
	}
	if mh.heap[mh.leftChildIndex(index)] < mh.heap[mh.rightChildIndex(index)] {
		return mh.leftChildIndex(index)
	} else {
		return mh.rightChildIndex(index)

	}

}

func (mh *miniHeap) heapifyUp(index int) {
	for mh.parentIndex(index) > 0 {
		if mh.heap[index] < mh.heap[mh.parentIndex(index)] {
			mh.swap(index, mh.parentIndex(index))
		}
		index = mh.parentIndex(index)
	}
}

func (mh *miniHeap) heapifyDown(index int) {
	for mh.leftChildIndex(index) <= mh.size {
		minChildIndex := mh.minChildIndex(index)
		if mh.heap[index] > mh.heap[minChildIndex] {
			mh.swap(index, minChildIndex)
		}
		index = minChildIndex
	}
}

func (mh *miniHeap) push(value int) {
	mh.heap = append(mh.heap, value)
	mh.size++
	mh.heapifyUp(mh.size)
}

func (mh *miniHeap) pop() *int {
	if (len(mh.heap)) == 1 {
		return nil
	}

	root := mh.heap[1]
	data := mh.heap[len(mh.heap)-1]
	mh.heap = mh.heap[:len(mh.heap)-1]

	if len(mh.heap) == 1 {
		return &root
	}

	mh.heap[1] = data
	mh.size--
	mh.heapifyDown(1)

	return &root
}

func main() {
	h := miniHeap{heap: []int{-1}}
	h.push(5)
	h.push(6)
	h.push(2)
	h.push(9)
	h.push(13)
	h.push(11)
	h.push(1)
	fmt.Println(h)
	fmt.Println(*h.pop())
	fmt.Println(h)
}
