package main

import (
	"math/rand"
	"slices"
)

//	Starting With a Binary Search Tree

//const LEAF_LEN int = 2

type AVL struct {
	root *AVL_node
}

type AVL_node struct {
	left   *AVL_node
	right  *AVL_node
	parent *AVL_node
	value  int
	height int
}

func createAVL_node(val int) *AVL_node {
	newAVL_node := &AVL_node{left: nil, right: nil, parent: nil, value: val, height: 1}
	return newAVL_node
}

func createAVL() AVL {
	tree := AVL{root: nil}
	return tree
}

func getHeight(root *AVL_node) int {
	if root == nil {
		return 0
	} else {
		return root.height
	}
}

func balanceFactor(root *AVL_node) int {
	return getHeight(root.right) - getHeight(root.left)
}

func rotateRight(root *AVL_node) *AVL_node {
	newroot := root.left
	root.left = root.left.right
	if root.left != nil {
		root.left.parent = root
	}
	newroot.right = root

	newroot.parent = root.parent
	root.parent = newroot

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	return newroot
}

func rotateLeft(root *AVL_node) *AVL_node {
	newroot := root.right
	root.right = root.right.left
	if root.right != nil {
		root.right.parent = root
	}
	newroot.left = root

	newroot.parent = root.parent
	root.parent = newroot

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	return newroot
}

func reBalance(tree *AVL, root *AVL_node) *AVL_node {

	// 4 cases LL, RR, LR, RL
	// LL
	if balanceFactor(root) < -1 && balanceFactor(root.left) < 0 {
		return rotateRight(root)
	}

	// RR
	if balanceFactor(root) > 1 && balanceFactor(root.right) > 0 {
		return rotateLeft(root)
	}

	// LR
	if balanceFactor(root) < -1 && balanceFactor(root.left) > 0 {
		root.left = rotateLeft(root.left)
		newroot := rotateRight(root)
		newroot.height = 1 + max(getHeight(newroot.right), getHeight(newroot.left))
		return newroot
	}

	// RL
	if balanceFactor(root) > 1 && balanceFactor(root.right) < 0 {
		root.right = rotateRight(root.right)
		newroot := rotateLeft(root)
		newroot.height = 1 + max(getHeight(newroot.right), getHeight(newroot.left))
		return newroot
	}

	return nil
}

func insertAVL(tree *AVL, root *AVL_node, n *AVL_node) {
	if tree.root == nil {
		tree.root = n
		return
	}

	if n.value == root.value || n.value == 0 {
		return
	}

	if n.value < root.value {
		if root.left != nil {
			insertAVL(tree, root.left, n)
		} else {
			n.parent = root
			root.left = n
		}
	} else {
		if root.right != nil {
			insertAVL(tree, root.right, n)
		} else {
			n.parent = root
			root.right = n
		}
	}

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	newparent := root.parent
	newroot := reBalance(tree, root)
	if newroot != nil {
		if newroot.left == tree.root || newroot.right == tree.root {
			tree.root = newroot
			newroot.parent = nil
		} else if newparent != nil {
			if newparent.left == root {
				newparent.left = newroot
			} else if newparent.right == root {
				newparent.right = newroot
			}
		}
	}
}

func nextInorderAVL(root *AVL_node) *AVL_node {
	root = root.right
	for root != nil && root.left != nil {
		root = root.left
	}

	return root
}

func deleteAVL(tree *AVL, root *AVL_node, nval int) *AVL_node {
	if tree.root == nil || root == nil {
		return nil
	}

	if nval == 0 {
		return nil
	}

	if nval == root.value {

		// Cases when root has 0 children or
		var newroot *AVL_node = nil
		if root.left == nil {
			if tree.root == root {
				tree.root = root.right
				tree.root.parent = nil
			}
			if root.right != nil {
				root.right.parent = root.parent
			}
			return root.right

		} else if root.right == nil {
			if tree.root == root {
				tree.root = root.left
				tree.root.parent = nil
			}
			if root.left != nil {
				root.left.parent = root.parent
			}
			return root.left

		} else {
			// If both left and right have children
			newroot = nextInorderAVL(root)
			root.value = newroot.value
			root.right = deleteAVL(tree, root.right, newroot.value)

		}

	} else if nval < root.value {
		root.left = deleteAVL(tree, root.left, nval)
	} else if nval > root.value {
		root.right = deleteAVL(tree, root.right, nval)
	}

	root.height = 1 + max(getHeight(root.right), getHeight(root.left))
	newroot := reBalance(tree, root)
	if newroot != nil {
		if newroot.left == tree.root || newroot.right == tree.root {
			tree.root = newroot
			newroot.parent = nil
		}
		return newroot
	}
	return root
}

/*
func inorderPrint(root *BST_node) {
	if root != nil {
		inorderPrint(root.left)
		fmt.Print(root.value, " ")
		if root.left == nil {
			fmt.Print("nil ")
		} else {
			fmt.Print(root.left.value, " ")
		}
		if root.right == nil {
			fmt.Print("nil | ")
		} else {
			fmt.Print(root.right.value, " | ")
		}

		inorderPrint(root.right)
	}
}

func inorderCheck(root *BST_node) {
	if root != nil {
		inorderCheck(root.left)
		if root.left != nil {
			if root.left.value >= root.value {
				panic("Left Child Error")
			}
		}
		if root.right != nil {
			if root.right.value <= root.value {
				panic("Right Child Error")
			}
		}

		inorderCheck(root.right)
	}
}
*/

const RAND_RANGEAVL = 100000

func insertSortArray(ary *[]int, val int) {
	//func Insert[S ~[]E, E any](s S, i int, v ...E) S
	if len(*ary) == 0 {
		(*ary) = append((*ary), val)
		return
	}

	for i := 0; i < len(*ary); i++ {
		if val == (*ary)[i] {
			return
		}
		if val < (*ary)[i] {
			(*ary) = slices.Insert(*ary, i, val)
			return
		}
	}

	(*ary) = append((*ary), val)
}

func randomInsertAVL(tree *AVL, ary *[]int, itr int) {
	for range itr {
		rnum := rand.Intn(RAND_RANGEAVL) + 1
		newAVL_node := createAVL_node(rnum)
		insertSortArray(ary, rnum)
		insertAVL(tree, tree.root, newAVL_node)
		inorderCheck(tree.root)
	}
}

func randomDeleteAVL(tree *AVL, ary *[]int, itr int) {
	for range itr {
		if len(*ary) > 0 {
			rnum := rand.Intn(len(*ary))
			delval := (*ary)[rnum]
			(*ary) = slices.Delete(*ary, rnum, rnum+1)
			deleteAVL(tree, tree.root, delval)
			inorderCheck(tree.root)
		}
	}
}
