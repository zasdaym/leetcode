//lint:file-ignore U1000 leetcode
package maxsubarray

import "math"

func maxSubArray(nums []int) int {
	maxPrice := 0
	result := -math.MaxInt64

	for _, num := range nums {
		if maxPrice+num > num {
			maxPrice += num
		} else {
			maxPrice = num
		}

		if maxPrice > result {
			result = maxPrice
		}
	}

	return result
}
