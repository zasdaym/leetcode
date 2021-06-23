package diameterofbinarytree

type treeNode struct {
	val         int
	left, right *treeNode
}

func diameterOfBinaryTree(root *treeNode) int {
	_, diameter := helper(root)
	return diameter
}

func helper(root *treeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftHeight, leftDiameter := 0, 0
	rightHeight, rightDiameter := 0, 0

	if root.left != nil {
		leftHeight, leftDiameter = helper(root.left)
	}

	if root.right != nil {
		rightHeight, rightDiameter = helper(root.right)
	}

	height := max(leftHeight, rightHeight) + 1
	currDiameter := leftHeight + rightHeight
	diameter := max(currDiameter, leftDiameter, rightDiameter)
	return height, diameter
}

func max(num int, nums ...int) int {
	result := num
	for _, n := range nums {
		if n > result {
			result = n
		}
	}
	return result
}
