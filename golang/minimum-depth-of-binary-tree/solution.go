package minimumdepthofbinarytree

import (
	"container/list"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func minDepth(root *treeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		depth++
		n := queue.Len()
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*treeNode)
			if node.left == nil && node.right == nil {
				return depth
			}
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
	}

	return depth
}
