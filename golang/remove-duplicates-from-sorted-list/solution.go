//lint:file-ignore U1000 leetcode
package deleteduplicates

type listNode struct {
	val  int
	next *listNode
}

func deleteDuplicates(head *listNode) *listNode {
	curr := head
	for curr != nil && curr.next != nil {
		if curr.val == curr.next.val {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
	return head
}
