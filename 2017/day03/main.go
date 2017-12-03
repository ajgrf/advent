package main

import (
	"fmt"
	"math"
)

const addr = 325489

// pos represents the coordinates of a square on the grid.
type pos struct {
	x, y int
}

// coordinates finds the Cartesian coordinates of the given address, if 1 were
// located at (0, 0).
func coordinates(addr int) pos {
	var x, y int

	if addr < 1 {
		panic("coordinates: Memory addresses must be >= 1")
	}

	// The largest number which is still smaller than addr when squared
	largestSquare := int(math.Sqrt(float64(addr)))

	// diff is the remaining spaces we need to go past the largestSquare
	diff := addr - (largestSquare * largestSquare)
	// diffX and diffY break it down into up/down & left/right movements
	diffX, diffY := 0, 0
	if diff != 0 {
		diffX = min(1, largestSquare-diff+2)
		diffY = min(diff-1, largestSquare)
	}

	// Add or subtract diffX & diffY with the coordinate for largestSquare,
	// depending on whether it's an even square near a top left corner in the
	// spiral, or an odd square near a bottom right corner.
	if largestSquare%2 == 0 {
		x = -largestSquare/2 + 1 - diffX
		y = largestSquare/2 - diffY
	} else {
		x = largestSquare/2 + diffX
		y = -largestSquare/2 + diffY
	}

	return pos{x, y}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
