class TreeNode:
    def __init__(self, val: int, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right

def has_path_sum(root: TreeNode, target: int) -> bool:
    if not root:
        return False

    if not (root.left or root.right):
        return root.val == target
    
    target -= root.val
    return has_path_sum(root.left, target) or has_path_sum(root.right, target)