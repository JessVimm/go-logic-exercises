package puzzles01

import (
	"fmt"
)

func CallPuzzles01() {
	fmt.Println("Puzzle Turning Rectangle Into Squares")
	squares := SquaresInRect(2, 3)
	fmt.Println("Result is:", squares)
}

func SquaresInRect(lng int, wdth int) []int {
	minSize, numSquares := 0, 0
	var sizeOfEachSquare []int

	if lng < wdth {
		minSize = lng
	} else if wdth < lng {
		minSize = wdth
	} else {
		return nil
	}

	numSquares = lng * wdth

	generateSqrs(&minSize, &numSquares, &sizeOfEachSquare)

	return sizeOfEachSquare
}

func generateSqrs(minSz *int, nSqrs *int, sizeOfSqrs *[]int) {
	if *nSqrs < 1 {
		return
	}
	*sizeOfSqrs = append(*sizeOfSqrs, *minSz)
	*nSqrs = *nSqrs - ((*minSz) * (*minSz))
	nextSize := *nSqrs / (*minSz)

	if nextSize >= *minSz {
		generateSqrs(minSz, nSqrs, sizeOfSqrs)
	} else {
		generateSqrs(&nextSize, nSqrs, sizeOfSqrs)
	}
}
