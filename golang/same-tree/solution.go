package sametree

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func isSameTree(p, q *treeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.val != q.val {
		return false
	}

	return isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
}
