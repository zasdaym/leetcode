//lint:file-ignore U1000 leetcode
package reverselist

type listNode struct {
	val  int
	next *listNode
}

func reverseList(head *listNode) *listNode {
	var prev *listNode
	for head != nil {
		head.next, head, prev = prev, head.next, head
	}
	return prev
}
