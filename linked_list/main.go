package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head *node
}

func (l *linkedList) append(data int) {
	newNode := node{data: data}
	if l.head == nil {
		l.head = &newNode
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = &newNode
}

func (l *linkedList) insert(data int) {
	newNode := node{data: data}
	newNode.next = l.head
	l.head = &newNode
}

func (l *linkedList) print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *linkedList) remove(data int) {
	currentNode := l.head
	if currentNode != nil && currentNode.data == data {
		l.head = currentNode.next
		currentNode = nil
		return
	}

	var previousNode *node
	for currentNode != nil && currentNode.data != data {
		previousNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return
	}

	previousNode.next = currentNode.next
	currentNode = nil
}

func (l *linkedList) reverseIterative() {
	var previousNode *node
	currentNode := l.head
	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}
	l.head = previousNode
}

func (l *linkedList) reverseRecursive() {
	var _reverseRecursive func(currentNode, previousNode *node) *node
	_reverseRecursive = func(currentNode, previousNode *node) *node {
		if currentNode == nil {
			return previousNode
		}
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
		return _reverseRecursive(currentNode, previousNode)
	}
	l.head = _reverseRecursive(l.head, nil)
}

func (l *linkedList) reverseEven() {
	var _reverseEven func(head, previousNode *node) *node
	_reverseEven = func(head, previousNode *node) *node {
		if head == nil {
			return nil
		}
		currentNode := head
		for currentNode != nil && currentNode.data%2 == 0 {
			nextNode := currentNode.next
			currentNode.next = previousNode

			previousNode = currentNode
			currentNode = nextNode
		}

		if currentNode != head {
			head.next = currentNode
			_reverseEven(currentNode, nil)
			return previousNode
		} else {
			head.next = _reverseEven(head.next, head)
			return head
		}
	}
	l.head = _reverseEven(l.head, nil)
}

func main() {
	l := linkedList{}
	l.append(1)
	l.append(2)
	l.append(2)
	l.insert(0)
	fmt.Println(l.head.data)
	fmt.Println(l.head.next.data)
	fmt.Println(l.head.next.next.data)
	fmt.Println(l.head.next.next.next.data)
	l.print()
	l.remove(2)
	fmt.Println("reversed ite")
	l.reverseIterative()
	l.print()
	fmt.Println("reversed rec")
	l.reverseRecursive()
	l.print()

	l.append(1)
	l.append(3)
	l.append(5)
	l.append(2)
	l.append(4)
	l.append(6)

	l.append(1)
	l.append(3)
	l.append(5)
	l.reverseEven()
	fmt.Println("-----rev even")
	l.print()
}
