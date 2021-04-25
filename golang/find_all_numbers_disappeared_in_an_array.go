//lint:file-ignore U1000 leetcode
package golang

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}

	var result []int
	for i, num := range nums {
		if num < 0 {
			continue
		}
		result = append(result, i+1)
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
