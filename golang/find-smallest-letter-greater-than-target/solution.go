package nextgreatestletter

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)

	for left < right {
		middle := left + (right-left)/2
		if letters[middle] <= target {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return letters[left%len(letters)]
}
