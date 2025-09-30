package avl

import "fmt"

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
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

		root.height = a.height(root)

		// balance
		a.balance()

		return root
	}

	a.root = travel(a.root, value)

}

// rotate left(avl node) -> avl node
// rotate right(avl node) -> avl node
// balance method replace print with the method
// change the return type of balance method to avl node
// and assign the balance return to the insert method

func (a *AVLTree) balance() {
	if a.isLeftHeavy(a.root) {
		if a.balanceFactor(a.root.left) < 0 {
			fmt.Println("Left rotate ", a.root.right.value)
		}
		fmt.Println("Right rotate", a.root.value)
	} else if a.isRightHeavy(a.root) {
		if a.balanceFactor(a.root.right) > 0 {
			fmt.Println("Right rotate ", a.root.right.value)
		}
		fmt.Println("Left rotate", a.root.value)
	}
}

func (a *AVLTree) height(node *Node) int {
	if node == nil {
		return -1
	}

	left := a.height(node.left)
	right := a.height(node.right)

	var max int

	if left > right {
		max = left
	} else {
		max = right
	}

	return 1 + max
}

func (a *AVLTree) isLeftHeavy(node *Node) bool {
	return a.balanceFactor(node) > 1
}

func (a *AVLTree) isRightHeavy(node *Node) bool {
	return a.balanceFactor(node) < -1
}

func (a *AVLTree) balanceFactor(node *Node) int {
	return a.height(node.left) - a.height(node.right)
}

func Init() {
	fmt.Println("AVL")

	tree := AVLTree{}

	fmt.Println("Insert")

	// tree.insert(7)
	// tree.insert(4)
	// tree.insert(9)
	// tree.insert(1)
	// tree.insert(6)
	// tree.insert(8)
	// tree.insert(10)
	// tree.insert(30)
	// tree.insert(40)

	// tree.insert(12)
	// tree.insert(3)
	// tree.insert(9)

	tree.insert(10)
	tree.insert(30)
	tree.insert(20)

	// tree.insert(4)
	fmt.Println("WASD")
}
