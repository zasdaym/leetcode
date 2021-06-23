package diameterofbinarytree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *treeNode
		expected int
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name: "unbalanced tree",
			root: &treeNode{
				val: 1,
				left: &treeNode{
					val: 2,
					left: &treeNode{
						val: 4,
					},
					right: &treeNode{
						val: 5,
					},
				},
				right: &treeNode{
					val: 3,
				},
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tc.root)
			if tc.expected != got {
				t.Errorf("want %d got %d", tc.expected, got)
			}
		})
	}
}
