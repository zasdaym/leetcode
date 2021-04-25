//lint:file-ignore U1000 leetcode
package golang

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	result := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		profit := price - minPrice
		if profit > result {
			result = profit
		}
	}

	return result
}
