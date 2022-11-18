// Given a string s, find the length of the longest substring without repeating characters.
package strings02

import (
	"fmt"
)

func StringsProblemTwo() {
	var str string = "cinturita"
	numChars := lengthOfLongestSubstring(str)
	fmt.Printf("Longest substring has %v characters...", numChars)
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 1

	if len(s) > 0 {
		for base := 0; base < len(s); base++ {
			var charsSeen []string
			charsSeen = append(charsSeen, string(s[base]))
			baseLength := 1
			foundSmilarChar := false

			if len(s)-base <= maxLength {
				break
			}

			for p := base + 1; p < len(s); p++ {
				// Check if current char has aleready been seen
				for _, c := range charsSeen {
					if c == string(s[p]) {
						// Look elsewhere
						foundSmilarChar = true
						break
					}
				}

				if foundSmilarChar {
					break
				} else {
					baseLength++
					if baseLength > maxLength {
						maxLength = baseLength
					}
					charsSeen = append(charsSeen, string(s[p]))
				}
			}
		}
	} else {
		maxLength = 0
	}

	return maxLength
}
