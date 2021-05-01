from typing import List


def next_greatest_letter(letters: List[int], target: str) -> str:
    left, right = 0, len(letters)
    while left < right:
        middle = (left + right) // 2
        if letters[middle] <= target:
            left = middle + 1
        else:
            right = middle
    return letters[left % len(letters)]

assert next_greatest_letter(["a", "b", "c", "d", "f", "g"], "e") == "f"
assert next_greatest_letter(["c", "d", "k", "j", "f", "g"], "p") == "c"
