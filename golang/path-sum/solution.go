package pathsum

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func hasPathSum(root *treeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.left == nil && root.right == nil {
		return root.val == target
	}

	target -= root.val

	return hasPathSum(root.left, target) || hasPathSum(root.right, target)
}
