class Node:
    def __init__(self, val: int, next=None):
        self.val = val
        self.next = next
    
def has_cycle(head: Node) -> bool:
    if not head or not head.next:
        return False
    
    slow = fast = head
    while fast.next and fast.next.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    
    return False