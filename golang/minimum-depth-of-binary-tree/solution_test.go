package minimumdepthofbinarytree

import "testing"

func TestMinDepth(t *testing.T) {
	testCases := []struct {
		name     string
		root     *treeNode
		expected int
	}{
		{
			name:     "empty node",
			root:     nil,
			expected: 0,
		},
		{
			name: "non-empty node",
			root: &treeNode{
				val: 3,
				left: &treeNode{
					val: 9,
				},
				right: &treeNode{
					val: 20,
					left: &treeNode{
						val: 15,
					},
					right: &treeNode{
						val: 7,
					},
				},
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minDepth(tc.root)
			if tc.expected != got {
				t.Errorf("want %d got %d", tc.expected, got)
			}
		})
	}
}
