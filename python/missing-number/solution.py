from typing import List


def missing_number(nums: List[int]) -> int:
    total = (len(nums) + 1) * len(nums) // 2
    for num in nums:
        total -= num
    return total
