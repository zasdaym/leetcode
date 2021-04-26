//lint:file-ignore U1000 leetcode
package missingnumber

func missingNumber(nums []int) int {
	n := len(nums)
	total := (n + 1) * n / 2
	for _, num := range nums {
		total -= num
	}
	return total
}
