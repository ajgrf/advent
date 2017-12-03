package main

import (
	"fmt"
	"math"
)

const addr = 325489

// coordinates finds the Cartesian coordinates of the given address, if 1 were
// located at (0, 0).
func coordinates(addr int) (x, y int) {
	largestSquare := int(math.Sqrt(float64(addr)))
	diff := addr - (largestSquare * largestSquare)

	var diffX, diffY int
	if corner := largestSquare + 1; diff == 0 {
		diffX = 0
		diffY = 0
	} else if diff <= corner {
		diffX = 1
		diffY = diff - 1
	} else if diff >= corner {
		diffX = -(diff - largestSquare - 2)
		diffY = largestSquare
	}

	if largestSquare%2 == 0 {
		x = -largestSquare/2 + 1 - diffX
		y = largestSquare/2 - diffY
	} else {
		x = largestSquare/2 + diffX
		y = -x + diffY
	}
	return
}

func distance(x, y int) int {
	return abs(x) + abs(y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	// Part 1
	fmt.Println(distance(coordinates(addr)))
}
