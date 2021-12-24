package main

import (
	// External imports
	"fmt"

	"./utils"
)

type avl struct {
	left   *avl
	right  *avl
	key    int
	height int
}

func (node *avl) getHeight() int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}

func (node *avl) getBF() int {
	if node == nil {
		return 0
	}

	lh := node.left.getHeight()
	rh := node.right.getHeight()

	return rh - lh
}

func (node *avl) getMinNode() *avl {
	if node == nil {
		return node
	} else if node.left == nil {
		return node
	} else {
		return node.left.getMinNode()
	}
}

func (node *avl) getMaxNode() *avl {
	if node == nil {
		return node
	} else if node.right == nil {
		return node
	} else {
		return node.left.getMaxNode()
	}
}

// Return new root
func (node *avl) rotateLeft() *avl {
	temp := node.right
	node.right = node.right.left
	temp.left = node

	// Adjust height
	node.height = utils.GetMax(node.left.getHeight(), node.right.getHeight()) + 1
	temp.height = utils.GetMax(temp.left.getHeight(), temp.right.getHeight()) + 1

	return temp
}

// Return new root
func (node *avl) rotateRight() *avl {
	temp := node.left
	node.left = node.left.right
	temp.right = node

	fmt.Println("left", temp.left)
	fmt.Println("right", temp.right)

	// Adjust height
	node.height = utils.GetMax(node.left.getHeight(), node.right.getHeight()) + 1
	temp.height = utils.GetMax(temp.left.getHeight(), temp.right.getHeight()) + 1

	fmt.Println("height", temp.height)

	return temp
}

// Return new root (if rotated)
func (root *avl) insert(key int) *avl {
	// Inserting
	if key > root.key {
		if root.right == nil {
			root.right = &avl{nil, nil, key, 0}
		} else {
			root.right.insert(key)
		}
	} else { // key < root.val
		if root.left == nil {
			root.left = &avl{nil, nil, key, 0}
		} else {
			root.left.insert(key)
		}
	}

	// Updating Height
	root.height = utils.GetMax(root.left.getHeight(), root.right.getHeight()) + 1

	// Balancing
	bf := root.getBF()

	// Left Heavy
	if bf < -1 {
		// Single rotation
		if root.left.getBF() <= 0 {
			return root.rotateRight()
		} else {
			// Need double rotation
			temp := root.left.rotateLeft()
			return temp.rotateRight()
		}
	} else if bf > 1 {
		// Single rotation
		if root.right.getBF() >= 1 {
			return root.rotateLeft()
		} else {
			// Need double rotation
			temp := root.right.rotateRight()
			return temp.rotateLeft()
		}
	}

	// No rotation
	return root
}

// Precondition: the key must be in the tree
func (root *avl) delete(key int) *avl {
	if root == nil {
		return root
	}

	// BST Deletion
	if root.key < key {
		root.left.delete(key)
	} else if root.key > key {
		root.right.delete(key)
	} else {
		// If node only has a child
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		}
		if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}

		succ := root.right.getMinNode()
		root.right = root.right.delete(succ.key)
	}

	// Check if root has one child
	if root == nil {
		return root
	}

	// TODO: Adjusting height

	// TODO: Balancing the three

	// Placeholder
	return nil
}

func main() {
	root := &avl{nil, nil, 10, 0}
	root = root.insert(5)
	root = root.insert(1)
	fmt.Println(root)
}
