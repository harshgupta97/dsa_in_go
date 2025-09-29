package avl

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type AVLTree struct {
	root *Node
}

func (a *AVLTree) insert(value int) {

	var travel func(root *Node, value int) *Node

	travel = func(root *Node, value int) *Node {
		if root == nil {
			return &Node{value: value}
		}

		// left
		if value < root.value {
			root.left = travel(root.left, value)
		}

		// right
		if value > root.value {
			root.right = travel(root.right, value)
		}

		// none
		if value == root.value {
			fmt.Println(value)
		}

		return root
	}

	a.root = travel(a.root, value)

}

func Init() {
	fmt.Println("AVL")

	tree := AVLTree{}

	fmt.Println("Insert")

	tree.insert(7)
	tree.insert(4)
	tree.insert(9)
	tree.insert(1)
	tree.insert(6)
	tree.insert(8)
	tree.insert(10)
}
