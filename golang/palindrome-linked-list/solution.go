//lint:file-ignore U1000 leetcode
package palindromelinkedlist

type listNode struct {
	val  int
	next *listNode
}

func isPalindrome(head *listNode) bool {
	fast, slow := head, head
	var rev *listNode
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		oldRev := rev
		rev, slow = slow, slow.next
		rev.next = oldRev
	}

	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		if slow.val != rev.val {
			return false
		}
		slow, rev = slow.next, rev.next
	}

	return true
}
