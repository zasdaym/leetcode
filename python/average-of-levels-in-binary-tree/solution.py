from typing import List, Deque
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def average_of_levels(root: TreeNode) -> List[float]:
    result: List[float] = []
    queue: Deque[TreeNode] = deque()

    if not root:
        return result

    queue.append(root)
    while queue:
        n = len(queue)
        sum = 0.0
        for i in range(n):
            node = queue.popleft()
            sum += node.val
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(sum / n)

    return result
