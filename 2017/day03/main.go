package main

import (
	"fmt"
	"math"
)

const addr = 325489

type pos struct {
	x, y int
}

// coordinates finds the Cartesian coordinates of the given address, if 1 were
// located at (0, 0).
func coordinates(addr int) pos {
	if addr < 1 {
		panic("coordinates: Memory addresses must be >= 1")
	}

	var x, y, diffX, diffY int

	largestSquare := int(math.Sqrt(float64(addr)))
	diff := addr - (largestSquare * largestSquare)

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
		y = -largestSquare/2 + diffY
	}

	return pos{x, y}
}

func distanceFromAccessPort(p pos) int {
	return abs(p.x) + abs(p.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func neighbors(p pos) []pos {
	return []pos{
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
	}
}

func main() {
	// Part 1
	fmt.Println(distanceFromAccessPort(coordinates(addr)))

	// Part 2
	var grid = map[pos]int{{0, 0}: 1}
	for i := 1; ; i++ {
		p := coordinates(i)
		for _, neighbor := range neighbors(p) {
			grid[p] += grid[neighbor]
		}

		if val := grid[p]; val > addr {
			fmt.Println(val)
			break
		}
	}
}
