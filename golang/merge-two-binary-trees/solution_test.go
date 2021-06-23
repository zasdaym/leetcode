package mergetwobinarytrees

import "testing"

func TestMergeTrees(t *testing.T) {
	testCases := []struct {
		name           string
		a, b, expected *treeNode
	}{
		{
			name: "one empty tree",
			a: &treeNode{
				val: 1,
				left: &treeNode{
					val: 3,
				},
			},
			expected: &treeNode{
				val: 1,
				left: &treeNode{
					val: 3,
				},
			},
		},
		{
			name: "both non-empty tree",
			a: &treeNode{
				val: 1,
				left: &treeNode{
					val: 3,
					left: &treeNode{
						val: 5,
					},
				},
				right: &treeNode{
					val: 2,
				},
			},
			b: &treeNode{
				val: 2,
				left: &treeNode{
					val: 1,
					right: &treeNode{
						val: 4,
					},
				},
				right: &treeNode{
					val: 3,
					right: &treeNode{
						val: 7,
					},
				},
			},
			expected: &treeNode{
				val: 3,
				left: &treeNode{
					val: 4,
					left: &treeNode{
						val: 5,
					},
					right: &treeNode{
						val: 4,
					},
				},
				right: &treeNode{
					val: 5,
					right: &treeNode{
						val: 7,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeTrees(tc.a, tc.b)
			if !isSameTree(tc.expected, got) {
				t.Errorf("merged tree is different than expected")
			}
		})
	}
}

// isSameTree checks whether given binary trees have the same structure and value.
func isSameTree(a, b *treeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return isSameTree(a.left, b.left) && isSameTree(a.right, b.right)
}
