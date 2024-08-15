package tree

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

type Tree struct {
	root *treeNode
}

func (T *Tree) insertUtil(node *treeNode, val int) {
	if node.value > val {
		if node.left == nil {
			node.left = T.CreateNode(val)
		} else {
			T.insertUtil(node.left, val)
		}
	} else if node.value < val {
		if node.right == nil {
			node.right = T.CreateNode(val)
		} else {
			T.insertUtil(node.right, val)
		}
	}
}
func inorder(node *treeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.value, " ")

	inorder(node.right)
}

func (T *Tree) CreateNode(val int) *treeNode {
	newNode := new(treeNode)
	newNode.value = val
	return newNode
}

func (T *Tree) Insert(val int) {
	if T.root == nil {
		T.root = T.CreateNode(val)
		return
	}
	T.insertUtil(T.root, val)
}

func (T *Tree) InOrderTraverse() {
	inorder(T.root)
	fmt.Print("\n")
}

func Create() *Tree {
	return new(Tree)
}
