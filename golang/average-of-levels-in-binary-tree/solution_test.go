package averageoflevelsinbinarytree

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	testCases := []struct {
		name     string
		root     *treeNode
		expected []float64
	}{
		{
			name:     "nil root",
			root:     nil,
			expected: []float64{},
		},
		{
			name: "three levels",
			root: &treeNode{
				val: 16,
				left: &treeNode{
					val:   10,
					left:  &treeNode{val: 6},
					right: &treeNode{val: 12},
				},
				right: &treeNode{
					val:   20,
					left:  &treeNode{val: 18},
					right: &treeNode{val: 24},
				},
			},
			expected: []float64{16, 15, 15},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := averageOfLevels(tc.root)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("want %v got %v", tc.expected, got)
			}
		})
	}
}
