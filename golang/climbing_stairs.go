//lint:file-ignore U1000 leetcode
package golang

func climbStairs(steps int) int {
	if steps < 3 {
		return steps
	}

	stepVariations := make([]int, steps+1)
	stepVariations[1] = 1
	stepVariations[2] = 2

	for i := 3; i <= steps; i++ {
		stepVariations[i] = stepVariations[i-1] + stepVariations[i-2]
	}
	return stepVariations[steps]
}
