package tree

func updateUtil(root *treeNode, oldVal int) {
	if root.value == oldVal {
		if root.left != nil {
			root = root.left
			return
		} else {
			root = root.right
			return
		}
	} else if root.left.value == oldVal {
		if root.left.left != nil {
			root.left = root.left.left
			return
		} else {
			root.left = root.left.right
			return
		}
	} else if root.right.value == oldVal {
		if root.right.left != nil {
			root.right = root.right.left
			return
		} else {
			root.right = root.right.right
			return
		}
	}
	if oldVal < root.value {
		updateUtil(root.left, oldVal)
	} else if oldVal > root.value {
		updateUtil(root.left, oldVal)
	}
}
