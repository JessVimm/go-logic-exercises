package strings01

import (
	"fmt"
	"strings"
)

func StringsProblemOne() {
	var firstWord string = "Hell"
	var secondWord string = "Hello#"

	if backspaceCompare(firstWord, secondWord) {
		fmt.Println("They are equal!!!")
	} else {
		fmt.Println("They are not equal...")
	}
}

func backspaceCompare(s string, t string) bool {

	str1 := generateString(s)
	str2 := generateString(t)

	if str1 == str2 {
		return true
	} else {
		return false
	}
}

func generateString(str string) string {
	var strBuilt strings.Builder

	catCount := 0

	// Iterate over string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '#' {
			catCount++
		} else {
			// Check cat count
			if catCount == 0 {
				// Update new string
				strBuilt.WriteByte(str[i])
			} else {
				catCount--
			}
		}
	}

	newString := strBuilt.String()

	// Reverse the string
	runeString := []rune(newString) // Converting string to runes

	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		// Swap the values
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}

	return string(runeString) // Return rune turned into string
}
