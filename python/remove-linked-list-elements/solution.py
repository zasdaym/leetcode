class ListNode:
    def __init__(self, val: int, next: 'ListNode' = None):
        self.val = val
        self.next = next


def remove_elements(head: ListNode, val: int) -> ListNode:
    if not head:
        return head

    head.next = remove_elements(head.next, val)
    if head.val == val:
        return head.next

    return head
