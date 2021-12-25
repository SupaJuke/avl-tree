package main

import (
	// External imports
	"fmt"
	"strconv"

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
func (root *avl) Insert(key int) *avl {
	// Inserting
	if key > root.key {
		if root.right == nil {
			root.right = &avl{nil, nil, key, 0}
		} else {
			root.right = root.right.Insert(key)
		}
	} else { // key < root.val
		if root.left == nil {
			root.left = &avl{nil, nil, key, 0}
		} else {
			root.left = root.left.Insert(key)
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
func (root *avl) Delete(key int) *avl {
	if root == nil {
		return root
	}

	// BST Deletion
	if key < root.key {
		root.left = root.left.Delete(key)
	} else if key > root.key {
		root.right = root.right.Delete(key)
	} else {
		// If node is a leaf or has only 1 child
		if root.left == nil {
			root = root.right
			return root
		} else if root.right == nil {
			root = root.left
			return root
		}

		succ := root.right.getMinNode()
		root.key = succ.key
		root.right = root.right.Delete(succ.key)
	}

	// Check if root has one child
	if root == nil {
		return root
	}

	// Adjusting height
	root.height = utils.GetMax(root.left.getHeight(), root.right.getHeight()) + 1

	// Balancing the three
	bf := root.getBF()

	// Left heavy
	if bf < -1 {
		// Need to rotate twice
		if root.left.getBF() > 0 {
			root.left = root.left.rotateLeft()
		}
		root = root.rotateRight()
	}

	// Right heavy
	if bf > 1 {
		// Need to rotate twice
		if root.right.getBF() < -1 {
			root.right = root.right.rotateLeft()
		}
		root = root.rotateLeft()
	}

	// Placeholder
	return root
}

func (root *avl) PrintPreOrder() string {
	if root == nil {
		return ""
	}
	if root.left == nil && root.right == nil {
		return strconv.Itoa(root.key)
	} else {
		str := ""
		str += strconv.Itoa(root.key) + " "
		str += root.left.PrintPreOrder() + " "
		str += root.right.PrintPreOrder()
		return str
	}
}

func (root *avl) PrintInOrder() string {
	if root == nil {
		return ""
	}
	if root.left == nil && root.right == nil {
		return strconv.Itoa(root.key)
	} else {
		str := ""
		// Left child
		if root.left != nil {
			str += root.left.PrintInOrder() + " "
		}

		// Node
		str += strconv.Itoa(root.key) + " "

		// Right child
		if root.right != nil {
			str += root.right.PrintInOrder()
		}

		return str
	}
}

func main() {
	root := &avl{nil, nil, 10, 0}
	root = root.Insert(5)
	root = root.Insert(12)
	root = root.Insert(11)
	root = root.Insert(13)
	// fmt.Println(root.PrintPreOrder())
	fmt.Println(root.PrintInOrder())
	root = root.Delete(5)
	fmt.Println(root.PrintInOrder())
	fmt.Println(root)
}
