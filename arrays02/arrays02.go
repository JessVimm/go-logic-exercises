package arrays02

import (
	"fmt"
)

func ArraysProblemTwo() {
	arr := []int{1, 1}
	area := waterArea(arr)
	fmt.Println("Area of the container is:", area)
	fmt.Println("Area with optimize version:", getMaxArea(arr))
}

func waterArea(arr []int) int {
	maxArea := 0
	// Iterate and calculate area min(arr[i],arr[j]) * (j - i)
	for i := 0; i < len(arr) - 1; i++ {
		for j := i+1; j < len(arr); j++ {
			hght := min(arr[i], arr[j])
			width := j-i 
			currArea := hght * width
			// Check for maxArea
			if currArea > maxArea {
				// Update
				maxArea = currArea
			}
		}
	}
	return maxArea
}

func min(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func getMaxArea(arr []int) int {
	// Initialize our pointers and maxArea var
	firstP, maxArea, secondP := 0, 0, len(arr) - 1
	// Traverse arr with pointers
	for i := 0; firstP < secondP; i++ {
		// Calculate the area
		hght := min(arr[firstP], arr[secondP])
		width := secondP - firstP
		currArea := hght * width

		// Update maxArea
		if currArea > maxArea {
			maxArea = currArea
		}

		if arr[firstP] <= arr[secondP] {
			firstP++
		} else {
			secondP--
		}
	}
	return maxArea
}
