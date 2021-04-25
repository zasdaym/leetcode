from typing import List


class NumArray:
    def __init__(self, nums: List[int]):
        sums = [0]
        for num in nums:
            sums.append(sums[-1] + num)
        self.sums = sums
        print(sums)

    def sum_range(self, left: int, right: int) -> int:
        return self.sums[right + 1] - self.sums[left]
