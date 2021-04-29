package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		name     string
		nums     []int
		target   int
		expected int
	}

	testCases := []testCase{
		{
			name:     "existent number",
			nums:     []int{3, 4, 1, 2, 5, 7, 12},
			target:   5,
			expected: 4,
		},
		{
			name:     "non-existent number",
			nums:     []int{3, 4, 1, 2, 5, 7, 12},
			target:   100,
			expected: -1,
		},
		{
			name:     "single number",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := binarySearch(tc.nums, tc.target)
			if tc.expected != got {
				t.Errorf("want %v got %v", tc.expected, got)
			}
		})
	}
}
