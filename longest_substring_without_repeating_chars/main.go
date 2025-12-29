package main

import "fmt"

func main() {
	s := "abcabcbacabcabb"
	fmt.Println(lengthOfLongestSubstring(s))
}

// abcdabcabcabcd

// Visualization:
// a b c d a b c a b c d
// 0 1 2 3 4 5 6 7 8 9 10 11

// The longest substring without repeating characters is "abcd" with length 4.

func lengthOfLongestSubstring(s string) int {
	last := make(map[rune]int)
	l, maxlen := 0, 0

	for r, ch := range s {
		if idx, ok := last[ch]; ok && idx >= l {
			l = idx + 1
		}

		last[ch] = r

		if r-l+1 > maxlen {
			maxlen = r - l + 1
		}
	}

	return maxlen
}
