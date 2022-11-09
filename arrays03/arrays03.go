// Trapping Water problem
package arrays03

import (
	"fmt"
)

func ArraysProblemThree() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	totalTrappedWater := getTrappedWater(arr)
	totalOptimalWater := getOptimalWater(arr)
	fmt.Println("Total trapped water: ", totalTrappedWater)
	fmt.Println("Total trapped water with optimal solution: ", totalOptimalWater)
}

func getTrappedWater(arr []int) int {
	totalWater := 0

	// Get water
	for i := 0; i < len(arr); i++ {
		maxLeft := getMaxLeft(i, arr)
		maxRight := getMaxRight(i, arr)

		waterAbove := getMin(maxLeft, maxRight) - arr[i]

		if waterAbove >= 0 {
			totalWater += waterAbove
		}
	}
	return totalWater
}

func getMaxLeft(idx int, arr []int) int {
	max := 0
	for i := idx - 1; i >= 0; i-- {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func getMaxRight(idx int, arr []int) int {
	max := 0
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func getMin(left int, right int) int {
	if left < right {
		return left
	}
	return right
}

func getOptimalWater(arr []int) int {
	totalWater, pL, maxLeft, maxRight, side := 0, 0, 0, 0, 0
	pR := len(arr) - 1

	for iter := 0; pL < 3; iter++ {
		// Check which pointer is smallest
		if arr[pL] <= arr[pR] {
			// Focus on left pointer
			side = 0
			workOnThisP(arr, &pL, &maxLeft, &totalWater, side)
		} else {
			// Focus on rigth pointer
			side = 1
			workOnThisP(arr, &pR, &maxRight, &totalWater, side)
		}
	}
	return totalWater
}

func workOnThisP(arr []int, pointer *int, max *int, water *int, side int) {
	// Check if we update pointer or calculate waterAbove
	if arr[*pointer] < *max {
		// Calculate above water
		waterAbove := *max - arr[*pointer]

		if waterAbove >= 0 {
			*water += waterAbove
		}
	} else {
		// Update smaller pointer's value
		*max = arr[*pointer]
	}
	// Update pointer position
	if side == 1 {
		// Right side moves left
		*pointer--
	} else {
		// Left side moves right
		*pointer++
	}
}
