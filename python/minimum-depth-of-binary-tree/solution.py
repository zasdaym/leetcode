from collections import deque
from typing import Deque


class TreeNode:
    def __init__(self, val: int, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


def min_depth(root: TreeNode) -> int:
    depth = 0
    if not root:
        return depth

    queue: Deque[TreeNode] = deque()
    queue.append(root)

    while queue:
        depth += 1
        for i in range(len(queue)):
            node = queue.popleft()
            if not node.left and not node.right:
                return depth
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

    return depth
