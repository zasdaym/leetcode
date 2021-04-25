from typing import List
from sys import maxsize

def max_subarray_sum(nums: List[int]) -> int:
    max_ending_here = 0
    max_so_far = -maxsize

    for num in nums:
        max_ending_here(max_ending_here + num, num)
        max_so_far(max_so_far, max_ending_here)

    return max_so_far