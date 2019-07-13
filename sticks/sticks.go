package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Stick is the line segment on a two-dimensional plane going from point 1 to
// point 2, inclusive.
type Stick struct {
	ID     int
	X1, Y1 float64
	X2, Y2 float64
}

// ParseStick takes a string in the format "n:x1,y1,x2,y2" and returns the
// corresponding Stick.
func ParseStick(s string) (Stick, error) {
	ss := strings.Split(s, ":")
	idStr, s := ss[0], ss[1]

	ss = strings.SplitN(s, ",", 4)
	if len(ss) != 4 {
		return Stick{}, errors.New("can't parse stick in expected format " +
			"(\"n:x1,y1,x2,y2\")")
	}

	id := mustParseInt(idStr)
	x1 := mustParseFloat(ss[0])
	y1 := mustParseFloat(ss[1])
	x2 := mustParseFloat(ss[2])
	y2 := mustParseFloat(ss[3])

	return Stick{id, x1, y1, x2, y2}, nil
}

// IsAbove returns whether or not any part of Stick l is directly above m,
// assuming they do not overlap.
func (l Stick) IsAbove(m Stick) bool {
	switch {
	case l.At(m.X1) >= m.Y1,
		l.At(m.X2) >= m.Y2,
		m.At(l.X1) <= l.Y1,
		m.At(l.X2) <= l.Y2:

		return true
	}
	return false
}

// At returns l's y value at the specified x.
func (l Stick) At(x float64) float64 {
	// make sure point 2 isn't left of 1
	if l.X2 < l.X1 {
		l.X1, l.Y1, l.X2, l.Y2 = l.X2, l.Y2, l.X1, l.Y1
	}

	if l.X1 <= x && x <= l.X2 {
		if math.IsInf(l.slope(), 0) {
			// stick is vertical, and since sticks don't overlap we can choose
			// any y value on this stick
			return l.Y1
		}
		return l.Y1 + (x-l.X1)*l.slope()
	}
	return math.NaN()
}

// slope finds the slope of the Stick, which may be infinite if the Stick is
// vertical.
func (l Stick) slope() float64 {
	return (l.Y2 - l.Y1) / (l.X2 - l.X1)
}

// UPDATE ME:
// StickSlice implements sort.Interface for []Stick, producing a safe order in
// which to remove them from above without disturbing the other sticks. All
// sticks are assumed to not overlap.
type StickSlice []Stick

func (p StickSlice) Sort() {
	for i := range p {
		sticks := p[i:]
		for j := range sticks {
			if sticks.CanRemove(j) {
				p[i], p[i+j] = p[i+j], p[i]
				break
			}
		}
	}
}

func (p StickSlice) CanRemove(i int) bool {
	for j := range p {
		if i == j {
			continue
		} else if p[j].IsAbove(p[i]) {
			return false
		}
	}
	return true
}

// mustParseInt returns the int represented by s, or else.
func mustParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic("mustParseInt: " + err.Error())
	}
	return int(n)
}

// mustParseFloat returns the float64 represented by s, or else.
func mustParseFloat(s string) float64 {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("mustParseFloat: " + err.Error())
	}
	return x
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(os.Stderr, x)
			os.Exit(1)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("no input")
	}
	n := mustParseInt(scanner.Text())

	sticks := make([]Stick, n)

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			st, err := ParseStick(scanner.Text())
			if err != nil {
				panic(err)
			}
			sticks[i] = st
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	StickSlice(sticks).Sort()
	sortedIDs := make([]string, n)
	for i := range sticks {
		sortedIDs[i] = strconv.FormatInt(int64(sticks[i].ID), 10)
	}
	fmt.Println(strings.Join(sortedIDs, ", "))
}
