package tree

import "fmt"

type treeNode struct {
	Value int
	Left  *treeNode
	Right *treeNode
}

type Tree struct {
	Root *treeNode
}

func (T *Tree) CreateNode(val int) *treeNode {
	newNode := new(treeNode)
	newNode.Value = val
	return newNode
}

func (T *Tree) insertUtil(node *treeNode, val int) {
	if node.Left == nil {
		node.Left = T.CreateNode(val)
	} else {
		T.insertUtil(node.Left, val)
	}
	if node.Right == nil {
		node.Right = T.CreateNode(val)
	} else {
		T.insertUtil(node.Right, val)
	}
}

func (T *Tree) Insert(val int) {
	if T.Root == nil {
		T.Root = T.CreateNode(val)
		return
	}
	T.insertUtil(T.Root, val)
}

func PTree() {
	t := new(Tree)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)
	fmt.Println(t.Root.Value, t.Root.Left.Value, t.Root.Right.Value, "B tree")
}
