package strings01

import (
	"fmt"
	"strings"
)

func StringsProblemOne() {
	var firstWord string = "bbbextm"
	var secondWord string = "bbb#extm"

	if backspaceCompare(firstWord, secondWord) {
		fmt.Println("They are equal!!!")
	} else {
		fmt.Println("They are not equal...")
	}

	if optimalComparison(firstWord, secondWord) {
		fmt.Println("They are equal with the second algorithm!")
	} else {
		fmt.Println("Not equal with second algorithm...")
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

func optimalComparison(s string, t string) bool {
	p1, p2 := len(s)-1, len(t)-1 // Initializing pointers
	counter1, counter2 := 0, 0 // Cat counter

	for i := 0; p1 >= 0 || p2 >= 0; i++ {
		// Count '#'
		p1 = countCats(&p1, &counter1, s)
		p2 = countCats(&p2, &counter2, t)

		// Compare if both chars are not equal
		if p1 >= 0 && p2 >= 0 {
			if s[p1] != t[p2] {
				return false
			}
		} else if (p1 >= 0 && p2 < 0) || (p2 >= 0 && p1 < 0) {
			return false
		}
		// Move left
		p1--
		p2--
	}
	// Were equal
	return true
}

func countCats(pointer *int, count *int, str string) int {
	
	if *pointer < 0 {
		return *pointer
	} 

	if str[*pointer] != '#' && *count == 0 {
		return *pointer
	} else if str[*pointer] != '#' && *count > 0 {
		*pointer--
		*count--
	} else {
		// Continue to move one left and decrement '#' count
		*pointer--
		*count++
	}
	// Repeat until there is 0 count
	return countCats(pointer, count, str)
}
