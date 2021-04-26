//lint:file-ignore U1000 leetcode
package middlenode

type listNode struct {
	val  int
	next *listNode
}

func middleNode(head *listNode) *listNode {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
