package arrays01

import (
	"fmt"
)

func ArraysProblemOne() {
	var arr []int = []int{3, 2, 4}
	target := 6
	fmt.Println("Solution with O(n^2):")
	idxs := findIdxs(arr, target)
	fmt.Println(idxs)
	idxs = getIdxs(arr, target)
	fmt.Println("Solution with O(n):")
	fmt.Println(idxs)
}

func findIdxs(arr []int, target int) []int {
	var idxsArr []int
	// Traverse the array
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// Check if values add up to target
			if (arr[i] + arr[j]) == target {
				idxsArr = append(idxsArr, i)
				idxsArr = append(idxsArr, j)
				return idxsArr
			}
		}
	}
	return idxsArr
}

func getIdxs(arr []int, target int) []int {
	var idxs []int
	var wantedValues = make(map[int]int)
	// Traverse the array
	for i:= 0; i < len(arr); i++ {
		keyValue, keyFlag := wantedValues[arr[i]]

		if keyFlag {
			idxs = append(idxs, keyValue)
			idxs = append(idxs, i)
			return idxs
		}

		// Calculate desired value
		wantedValue := target - arr[i]
		// Save in map
		wantedValues[wantedValue] = i
	} 

	return idxs
}
