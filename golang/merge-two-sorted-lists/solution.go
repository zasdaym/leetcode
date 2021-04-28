//lint:file-ignore U1000 leetcode
package mergetwolists

type listNode struct {
	val  int
	next *listNode
}

func mergeTwoLists(l1, l2 *listNode) *listNode {
	dummyHead := &listNode{}
	curr := dummyHead

	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			curr.next, l1 = l1, l1.next
		} else {
			curr.next, l2 = l2, l2.next
		}
		curr = curr.next
	}

	if l1 != nil {
		curr.next = l1
	} else {
		curr.next = l2
	}

	return dummyHead.next
}
