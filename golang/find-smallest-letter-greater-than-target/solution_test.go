package nextgreatestletter

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	testCases := []struct {
		name     string
		letters  []byte
		target   byte
		expected byte
	}{
		{"target in the middle", []byte{'a', 'b', 'c', 'e', 'f'}, 'd', 'e'},
		{"target is wrap around", []byte{'a', 'c', 'f', 'j'}, 'p', 'a'},
		{"all letters is bigger", []byte{'b', 'c', 'f', 'j'}, 'a', 'b'},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := nextGreatestLetter(tc.letters, tc.target)
			if tc.expected != got {
				t.Errorf("want %c got %c", tc.expected, got)
			}
		})
	}
}
