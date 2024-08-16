package tree

import "fmt"

func filetest() {
	fmt.Println("q file testing... ok ")
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findHeight(root *treeNode) int {
	if root == nil {
		return -1
	}
	l := findHeight(root.left)
	r := findHeight(root.right)
	return 1 + max(l, r)
}
