class TreeNode:
    def __init__(self, val: int, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right

def is_same_tree(first: TreeNode, second: TreeNode) -> bool:
    if not first and not second:
        return True

    if not first or not second:
        return False

    if first.val != second.val:
        return False

    return is_same_tree(first.left, second.left) and is_same_tree(first.right, second.right)