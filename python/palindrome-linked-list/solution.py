class ListNode:
    def __init__(self, val: int, next: 'ListNode' = None):
        self.val = val
        self.next = next


def is_palindrome(head: ListNode) -> bool:
    slow = fast = head
    rev = None
    while fast and fast.next:
        fast = fast.next.next
        rev, slow, rev.next = slow, slow.next, rev

    if fast:
        slow = slow.next

    while slow:
        if slow.val != rev.val:
            return False
        slow = slow.next
        rev = rev.next

    return True
