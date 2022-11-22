// Given a string s, find the length of the longest substring without repeating characters.
package strings02

import (
	"fmt"
)

func StringsProblemTwo() {
	var str string = "aaaaffruiknla"
	numChars := lengthOfLongestSubstring(str)
	fmt.Printf("Longest substring with brute force solution has %v number of characters...", numChars)
	fmt.Println()
	numChars = optimalSolution(str)
	fmt.Printf("Longest substring with the optimal solution has %v number of characters", numChars)
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

func optimalSolution(s string) int {
	// Defining windows limits, hash map of seen chars and max len
	startPoint, maxLen := 0, 0
	seenChars := make(map[string]int)

	if len(s) <= 1 {
		return len(s)
	} else {
		// Do until endPoint is at end of s
		for endPoint:= 0; endPoint < len(s); endPoint++ {
			// Have we seen this char?
			idx, charAlreadySeen := seenChars[string(s[endPoint])]

			if charAlreadySeen {
				// Check if idx was already part of curr substring
				if idx >= startPoint {
					// There is a repetition
					// Move start point
					startPoint = idx + 1
				} 
			}
			// Update seen position for this char
			seenChars[string(s[endPoint])] = endPoint 
			windowSize := (endPoint - startPoint) + 1

			// Update maxlen
			if windowSize > maxLen {
				maxLen = windowSize
			}
		}
		return maxLen
	}
}
