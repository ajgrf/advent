package main

import "testing"

var coordinateTests = []struct {
	addr, x, y int
}{
	{1, 0, 0},
	{2, 1, 0},
	{3, 1, 1},
	{4, 0, 1},
	{5, -1, 1},
	{6, -1, 0},
	{7, -1, -1},
	{8, 0, -1},
	{9, 1, -1},
	{10, 2, -1},
	{11, 2, 0},
	{12, 2, 1},
	{13, 2, 2},
	{14, 1, 2},
	{15, 0, 2},
	{16, -1, 2},
	{17, -2, 2},
	{18, -2, 1},
	{19, -2, 0},
	{20, -2, -1},
	{21, -2, -2},
	{22, -1, -2},
	{23, 0, -2},
	{24, 1, -2},
	{25, 2, -2},
	{26, 3, -2},
	{addr, -267, -285},
}

func TestCoordinates(t *testing.T) {
	for _, tt := range coordinateTests {
		p := coordinates(tt.addr)
		if p.x != tt.x || p.y != tt.y {
			t.Errorf("coordinates(%d) - expected=(%d,%d), got=(%d,%d)", tt.addr, tt.x, tt.y, p.x, p.y)
		}
	}
}
