package sametree

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		name     string
		first    *treeNode
		second   *treeNode
		expected bool
	}{
		{
			name: "same tree",
			first: &treeNode{
				val: 5,
				left: &treeNode{
					val: 2,
				},
				right: &treeNode{
					val: 10,
				},
			},
			second: &treeNode{
				val: 5,
				left: &treeNode{
					val: 2,
				},
				right: &treeNode{
					val: 10,
				},
			},
			expected: true,
		},
		{
			name: "different structure",
			first: &treeNode{
				val: 5,
				left: &treeNode{
					val: 2,
				},
				right: &treeNode{
					val: 10,
				},
			},
			second: &treeNode{
				val: 5,
				right: &treeNode{
					val: 10,
				},
			},
			expected: false,
		},
		{
			name: "different value",
			first: &treeNode{
				val: 5,
				left: &treeNode{
					val: 2,
				},
				right: &treeNode{
					val: 10,
				},
			},
			second: &treeNode{
				val: 5,
				left: &treeNode{
					val: 4,
				},
				right: &treeNode{
					val: 10,
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSameTree(tc.first, tc.second)
			if tc.expected != got {
				t.Errorf("want %v got %v", tc.expected, got)
			}
		})
	}
}
