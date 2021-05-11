package averageoflevelsinbinarytree

import "container/list"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func averageOfLevels(root *treeNode) []float64 {
	result := make([]float64, 0)

	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		n := queue.Len()
		total := 0
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*treeNode)
			total += node.val
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
		result = append(result, float64(total)/float64(n))
	}

	return result
}
