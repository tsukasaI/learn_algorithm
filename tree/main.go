package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func (b *binarySearchTree) insert(v int) {
	if b.root == nil {
		b.root = &node{value: v}
	}
	var _insert func(n *node, v int) *node
	_insert = func(n *node, v int) *node {
		if n == nil {
			return &node{value: v}
		}

		if v < n.value {
			n.left = _insert(n.left, v)
		} else {
			n.right = _insert(n.right, v)
		}
		return n
	}

	_insert(b.root, v)
}

// left, root, right
func (b *binarySearchTree) inOrder() {
	fmt.Println("inorder")
	var _inOrder func(n *node)
	_inOrder = func(n *node) {
		if n != nil {
			_inOrder(n.left)
			fmt.Println(n.value)
			_inOrder(n.right)
		}
	}
	_inOrder(b.root)
}

func (b *binarySearchTree) search(v int) bool {
	var _search func(n *node, v int) bool
	_search = func(n *node, v int) bool {
		if n == nil {
			return false
		}

		if n.value == v {
			return true
		} else if n.value > v {
			return _search(n.left, v)
		} else {
			return _search(n.right, v)
		}
	}
	return _search(b.root, v)
}

func (b *binarySearchTree) minValue(n *node) *node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (b *binarySearchTree) remove(v int) {
	var _remove func(n *node, v int) *node
	_remove = func(n *node, v int) *node {
		if n == nil {
			return n
		}

		if v < n.value {
			n.left = _remove(n.left, v)
		} else if v > n.value {
			n.right = _remove(n.right, v)
		} else {
			if n.left == nil {
				return n.right
			} else if n.right == nil {
				return n.left
			}
			temp := b.minValue(n.right)
			n.value = temp.value
			n.right = _remove(n.right, temp.value)
		}
		return n
	}
	_remove(b.root, v)
}

func main() {
	var bt binarySearchTree

	bt.insert(3)
	bt.insert(6)
	bt.insert(5)
	bt.insert(7)
	bt.insert(1)
	bt.insert(10)
	bt.insert(2)
	bt.inOrder()
	fmt.Println(bt.search(6))
	bt.remove(6)
	bt.inOrder()
}
