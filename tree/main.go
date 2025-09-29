package main

import (
	"fmt"
	"tree/avl"
)

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

func (t *Tree) equality(tree *Tree) bool {
	return t.equalitys(t.root, tree.root)
}

func (t *Tree) equalitys(first, second *Node) bool {
	if first == nil && second == nil {
		return true
	}

	if first != nil && second != nil {
		return first.value == second.value && t.equalitys(first.left, second.left) && t.equalitys(first.right, second.right)
	}

	return false
}

func (t *Tree) validateBST(node *Node) bool {
	var helper func(node *Node, min, max *int) bool

	helper = func(node *Node, min, max *int) bool {

		if node == nil {
			return true
		}

		if min != nil && node.value < *min {
			return false
		}

		if max != nil && node.value > *max {
			return false
		}

		return helper(node.left, min, &node.value) && helper(node.right, &node.value, max)
	}

	return helper(node, nil, nil)
}

func (t *Tree) nodeAtK(node *Node, distance int) []int {
	var nodes []int

	if node == nil {
		return nodes
	}

	if distance == 0 {
		nodes = append(nodes, node.value)
		return nodes
	}

	if node.left == nil && node.right == nil {
		return nodes
	}

	left := t.nodeAtK(node.left, distance-1)
	right := t.nodeAtK(node.right, distance-1)

	nodes = append(nodes, left...)
	nodes = append(nodes, right...)

	return nodes
}

func (t *Tree) levelOrderTraversal() {
	if t.root == nil {
		return
	}

	for i := 0; i <= t.height(t.root); i++ {
		fmt.Println(t.nodeAtK(t.root, i))
	}
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

	tree2 := &Tree{}

	tree2.insert(7)
	tree2.insert(4)
	tree2.insert(9)
	tree2.insert(1)
	tree2.insert(6)
	tree2.insert(8)
	tree2.insert(10)

	fmt.Println("Node at K")
	fmt.Println(tree.nodeAtK(tree.root, 2))

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

	fmt.Println("Equality")
	fmt.Println(tree.equality(tree2))

	fmt.Println("Valid BST")
	fmt.Println(tree.validateBST(tree.root))

	fmt.Println("Level Order Traversal")
	tree.levelOrderTraversal()

	fmt.Print("AVL")
	avl.Init()
}
