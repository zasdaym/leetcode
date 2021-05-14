package pathsum

import "testing"

func TestHasPathSum(t *testing.T) {
	testCases := []struct {
		name     string
		root     *treeNode
		target   int
		expected bool
	}{
		{
			name: "path sum exists",
			root: &treeNode{
				val: 5,
				left: &treeNode{
					val: 4,
					left: &treeNode{
						val: 11,
						left: &treeNode{
							val: 7,
						},
						right: &treeNode{
							val: 2,
						},
					},
				},
				right: &treeNode{
					val: 8,
					left: &treeNode{
						val: 13,
					},
				},
			},
			target:   22,
			expected: true,
		},
		{
			name:     "empty tree",
			root:     nil,
			target:   0,
			expected: false,
		},
		{
			name: "path sum doesn't exist",
			root: &treeNode{
				val: 1,
				right: &treeNode{
					val: 2,
				},
			},
			target:   1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := hasPathSum(tc.root, tc.target)
			if tc.expected != got {
				t.Errorf("want %v got %v", tc.expected, got)
			}
		})
	}
}
