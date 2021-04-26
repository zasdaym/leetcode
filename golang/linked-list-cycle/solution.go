//lint:file-ignore U1000 leetcode
package hascycle

type listNode struct {
	val  int
	next *listNode
}

func hasCycle(head *listNode) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}

	return false
}
