//lint:file-ignore U1000 leetcode
package removelinkedlistelements

type listNode struct {
	val  int
	next *listNode
}

func removeElements(head *listNode, val int) *listNode {
	if head == nil {
		return head
	}

	curr := head
	prev := &listNode{
		next: curr,
	}
	dummyHead := prev

	for curr != nil {
		if curr.val == val {
			prev.next = curr.next
		} else {
			prev = prev.next
		}
		curr = curr.next
	}

	return dummyHead.next
}
