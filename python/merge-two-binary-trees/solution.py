class TreeNode:
    def __init__(self, val: int, left: 'TreeNode', right: 'TreeNode') -> None:
        self.val = val
        self.left = left
        self.right = right


def merge_binary_trees(a: TreeNode, b: TreeNode) -> None:
    if not a:
        return b

    if not b:
        return a

    root = TreeNode(a.val + b.val)
    root.left = merge_binary_trees(a.left, b.left)
    root.right = merge_binary_trees(a.right, b.right)

    return root
