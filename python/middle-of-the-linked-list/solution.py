class ListNode:
    def __init__(self, val: int, next=None):
        self.val = val
        self.next = next

def middle_node(head: ListNode) -> ListNode:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow