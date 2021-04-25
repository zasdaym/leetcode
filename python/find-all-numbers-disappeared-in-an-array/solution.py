from typing import List


def find_disappeared_numbers(nums: List[int]) -> List[int]:
    n = len(nums)
    
    for i in range(n):
        index = abs(nums[i]) - 1
        if nums[index] > 0:
            nums[index] = -nums[index]
        
    result: List[int] = []
    for i in range(n):
        if nums[i] > 0:
            result.append(i + 1)
    
    return result