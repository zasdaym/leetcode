from typing import List, Set


def contains_duplicate(nums: List[int]) -> bool:
    seen: Set[int] = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False
