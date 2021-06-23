from typing import Tuple


class TreeNode:
    def __init__(self, val: int, left: 'TreeNode', right: 'TreeNode') -> None:
        self.val = val
        self.left = left
        self.right = right


def diameter_of_binary_tree(root: TreeNode) -> int:
    _, diameter = helper(root)
    return diameter


def helper(root: TreeNode) -> Tuple[int, int]:
    if not root:
        return 0, 0

    left_height, left_diameter = 0, 0
    right_height, right_diameter = 0, 0

    if root.left:
        left_height, left_diameter = helper(root.left)
    if root.right:
        right_height, right_diameter = helper(root.right)

    height = max(left_height, right_height) + 1
    curr_diameter = left_height + right_height
    diameter = max(curr_diameter, left_diameter, right_diameter)
