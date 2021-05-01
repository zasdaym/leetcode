from typing import List


def peak_index_in_mountain_array(nums: List) -> int:
    left, right = 0, len(nums) - 1

    while left < right:
        middle = (left + right) // 2
        if nums[middle] < nums[middle + 1]:
            left = middle + 1
        else:
            right = middle

    return left


assert peak_index_in_mountain_array([1, 3, 5, 25, 6, 4, 2])
