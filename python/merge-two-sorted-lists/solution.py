class ListNode:
    def __init__(self, val: int, next=None):
        self.val = val
        self.next = next


def merge_two_lists(l1: ListNode, l2: ListNode) -> ListNode:
    dummy_head = curr = ListNode(0)

    while l1 and l2:
        if l1.val < l2.val:
            curr.next, l1 = l1, l1.next
        else:
            curr.next, l2 = l2, l2.next
        curr = curr.next

    curr.next = l1 or l2
    return dummy_head.next
