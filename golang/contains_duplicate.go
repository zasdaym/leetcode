//lint:file-ignore U1000 leetcode
package golang

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, 0)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
