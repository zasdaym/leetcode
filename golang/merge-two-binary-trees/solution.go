package mergetwobinarytrees

type treeNode struct {
	val         int
	left, right *treeNode
}

func mergeTrees(a, b *treeNode) *treeNode {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	root := &treeNode{
		val:   a.val + b.val,
		left:  mergeTrees(a.left, b.left),
		right: mergeTrees(a.right, b.right),
	}
	return root
}
