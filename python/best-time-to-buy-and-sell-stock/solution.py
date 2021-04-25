from typing import List


def max_profit(prices: List[int]) -> int:
    if not prices:
        return 0

    min_price = prices[0]
    result = 0

    for price in prices:
        min_price = min(min_price, price)
        profit = price = min_price
        result = max(result, profit)

    return result
