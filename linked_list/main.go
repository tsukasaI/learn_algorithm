package main

import (
	"fmt"
)

type (
	node struct {
		data int
		next *node
	}
	linkedList struct {
		head *node
	}
)

type (
	doublyNode struct {
		data int
		next *doublyNode
		prev *doublyNode
	}
	doublyLinkedList struct {
		head *doublyNode
	}
)

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

func (d *doublyLinkedList) append(data int) {
	newNode := doublyNode{data: data}
	if d.head == nil {
		d.head = &newNode
		return
	}
	currentNode := d.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &newNode
	newNode.prev = currentNode
}

func (d *doublyLinkedList) insert(data int) {
	newNode := doublyNode{data: data}
	if d.head == nil {
		d.head = &newNode
		return
	}

	d.head.prev = &newNode
	newNode.next = d.head
	d.head = &newNode

}

func (d *doublyLinkedList) print() {
	currentNode := d.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (d *doublyLinkedList) remove(data int) {
	currentNode := d.head
	if currentNode != nil && currentNode.data == data {
		if currentNode.next == nil {
			currentNode = nil
			d.head = nil
			return
		} else {
			nextNode := currentNode.next
			nextNode.prev = nil
			d.head = nextNode
			return
		}
	}
	for currentNode != nil && currentNode.data != data {
		currentNode = currentNode.next

		if currentNode == nil {
			return
		}

		if currentNode.next == nil {
			prev := currentNode.prev
			prev.next = nil
			currentNode = nil
			return
		} else {
			nextNode := currentNode.next
			prevNode := currentNode.prev
			prevNode.next = nextNode
			nextNode.prev = prevNode
			currentNode = nil
			return
		}
	}
}

func (d *doublyLinkedList) reverseIterative() {
	var previousNode *doublyNode
	currentNode := d.head

	for currentNode != nil {
		previousNode = currentNode.prev
		currentNode.prev, currentNode.next = currentNode.next, currentNode.prev
		currentNode = currentNode.prev

	}

	if previousNode != nil {
		d.head = previousNode.prev
	}
}

func (d *doublyLinkedList) reverseRecursive() {
	var _reverseRecursive func(currentNode *doublyNode) *doublyNode
	_reverseRecursive = func(currentNode *doublyNode) *doublyNode {
		if currentNode == nil {
			return nil
		}

		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next

		if currentNode.prev == nil {
			return currentNode
		}

		return _reverseRecursive(currentNode.prev)
	}
	d.head = _reverseRecursive(d.head)
}

func (d *doublyLinkedList) bubbleSort() {
	if d.head == nil {
		return
	}

	currentNode := d.head
	for currentNode.next != nil {
		if currentNode.data > currentNode.next.data {
			currentNode.data, currentNode.next.data = currentNode.next.data, currentNode.data
		}
		currentNode = currentNode.next
	}
}

func main() {
	fmt.Println("----linkedlist-----")
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
	fmt.Println("----doublyLinkedlist-----")
	d := doublyLinkedList{}
	d.append(1)
	d.append(2)
	d.append(3)
	d.insert(0)
	d.print()
	fmt.Println("reversed ite")
	d.reverseIterative()
	d.print()
	fmt.Println("reversed rec")
	d.reverseRecursive()
	d.print()

	d.insert(5)
	fmt.Println("bubbleSort")
	d.bubbleSort()
	d.print()
}
