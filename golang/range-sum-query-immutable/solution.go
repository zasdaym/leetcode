//lint:file-ignore U1000 leetcode
package numarray

type numArray struct {
	sums []int
}

func newNumArray(nums []int) *numArray {
	sums := []int{0}
	for _, num := range nums {
		sums = append(sums, sums[len(sums)-1]+num)
	}
	return &numArray{sums: sums}
}

func (n *numArray) sumRange(left, right int) int {
	return n.sums[right+1] - n.sums[left]
}
