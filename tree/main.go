package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(value int) {
	newNode := &Node{value: value}

	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root

	for {
		if value < current.value {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
			continue
		}

		if value > current.value {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
			continue
		}

		if value == current.value {
			return
		}
	}
}

func (t *Tree) find(value int) bool {
	current := t.root
	for current != nil {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			return true
		}
	}
	return false
}

func (t *Tree) preOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d\n", node.value)
	t.preOrderTraversal(node.left)
	t.preOrderTraversal(node.right)
}

func (t *Tree) inOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	t.inOrderTraversal(node.left)
	fmt.Printf("%d\n", node.value)
	t.inOrderTraversal(node.right)
}

func (t *Tree) postOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	t.postOrderTraversal(node.left)
	t.postOrderTraversal(node.right)
	fmt.Printf("%d\n", node.value)
}

func (t *Tree) height(node *Node) int {
	if node == nil {
		// node (0) and edge (-1) count
		return -1
	}

	left := t.height(node.left)
	right := t.height(node.right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	fmt.Println("Binary Tree")

	tree := &Tree{}

	fmt.Println("Insert")

	tree.insert(7)
	tree.insert(4)
	tree.insert(9)
	tree.insert(1)
	tree.insert(6)
	tree.insert(8)
	tree.insert(10)

	fmt.Println("Find")

	fmt.Println(tree.find(5))
	fmt.Println(tree.find(6))

	fmt.Println("Traversal")

	fmt.Println("Pre")
	tree.preOrderTraversal(tree.root)

	fmt.Println("In")
	tree.inOrderTraversal(tree.root)

	fmt.Println("Post")
	tree.postOrderTraversal(tree.root)

	fmt.Println("Height")
	fmt.Println(tree.height(tree.root))
}
