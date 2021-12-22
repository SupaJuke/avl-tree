package main

import (
	// External imports
	"fmt"

	"./utils"
)

type node struct {
	left   *node
	right  *node
	key    int
	height int
}

func (node *node) getHeight() int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}

func (node *node) getBF() int {
	if node == nil {
		return 0
	}

	lh := node.left.getHeight()
	rh := node.right.getHeight()

	return rh - lh
}

func (node *node) getMinNode() *node {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	} else {
		return node.left.getMinNode()
	}
}

func (node *node) getMaxNode() *node {
	if node == nil {
		return nil
	} else if node.right == nil {
		return node
	} else {
		return node.left.getMaxNode()
	}
}

// Return new root
func (node *node) rotateLeft() *node {
	temp := node.right
	node.right = node.right.left
	temp.left = node

	// Adjust height
	node.height = utils.GetMax(node.left.getHeight(), node.right.getHeight()) + 1
	temp.height = utils.GetMax(temp.left.getHeight(), temp.right.getHeight()) + 1

	return temp
}

// Return new root
func (node *node) rotateRight() *node {
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
func (root *node) insert(key int) *node {
	// Inserting
	if key > root.key {
		if root.right == nil {
			root.right = &node{nil, nil, key, 0}
		} else {
			root.right.insert(key)
		}
	} else { // key < root.val
		if root.left == nil {
			root.left = &node{nil, nil, key, 0}
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

func (root *node) delete(key int) *node {
	// TODO: finish the method
	if key == root.key {

	}
}

func main() {
	root := &node{nil, nil, 10, 0}
	root = root.insert(5)
	root = root.insert(1)
	fmt.Println(root)
}
