class ListNode():
    def __init__(self, val: int = 0, next: 'ListNode' = None):
        self.val = val
        self.next = next


def reverse_list(head: ListNode) -> ListNode:
    prev = None
    while head:
        head.next, head, prev = prev, head.next, head
    return prev
