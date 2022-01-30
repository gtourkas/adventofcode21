package adventofcode21

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type entry struct {
	a point
	b point
}

func (e entry) isDiagonal() bool {
	return math.Abs(float64(e.a.x - e.b.x)) == math.Abs(float64(e.a.y - e.b.y))
}

func (e entry) getPointsFromUpperLeftToLowerRight() (point, point) {
	if e.a.y < e.b.y {
		return e.a, e.b
	}
	return e.a, e.b
}


func getPointFromStr(pointStr string) point {
	numbersStr := strings.Split(pointStr, ",")
	x, _ := strconv.Atoi(numbersStr[0])
	y, _ := strconv.Atoi(numbersStr[1])
	return point{
		x: x,
		y: y,
	}
}

func updateMaxFromPoint(p point, x *int, y *int)  {
	if p.x > *x {
		*x = p.x
	}
	if p.y > *y {
		*y = p.y
	}
}

func updateMaxFromEntry(e *entry, x *int, y *int)  {
	updateMaxFromPoint(e.a, x, y)
	updateMaxFromPoint(e.b, x, y)
}

type floorPoint struct {
	overlaps int
	exceededThreshold bool
}

func printFloor(floor [][]floorPoint) {
	for i := range floor {
		for j := range floor[i] {
			fmt.Printf("%d ", floor[i][j].overlaps)
		}
		fmt.Println("")
	}
}

func getPointsFromLine(line string) (point, point) {
	pointsStr := strings.Split(line, " -> ")
	a := getPointFromStr(pointsStr[0])
	b := getPointFromStr(pointsStr[1])
	return a, b
}

func parseLines(lines[] string, ignoreDiagonals bool) ([]entry, [][]floorPoint) {
	entries := make([]entry, len(lines))
	maxX := 0; maxY := 0
	for i, l := range lines {
		a, b := getPointsFromLine(l)
		e := &entries[i]
		e.a = a
		e.b = b
		if e.isDiagonal() && ignoreDiagonals {
			continue
		}
		updateMaxFromEntry(e, &maxX, &maxY)
	}

	floor := make([][]floorPoint, maxY+1)
	for i := range floor {
		floor[i] = make([]floorPoint, maxX+1)
	}

	return entries, floor
}

func countOverlappingInSingleLine(x1 int, y1 int, x2 int, y2 int, floor [][]floorPoint, threshold int, exceededThresholdCount *int) {
	xStep := 1
	if x1 > x2 {
		xStep = -1
	}
	yStep := 1
	if y1 > y2 {
		yStep = -1
	}

	x := x1
	for {
		y := y1
		for {
			pt := &floor[y][x]
			pt.overlaps += 1
			if !pt.exceededThreshold && pt.overlaps >= threshold {
				pt.exceededThreshold = true
				*exceededThresholdCount++
			}
			if (yStep == 1 && y + yStep > y2)  || (yStep == -1 && y + yStep < y2) {
				break
			}
			y += yStep
		}

		if (xStep == 1 && x + xStep > x2)  || (xStep == -1 && x + xStep < x2) {
			break
		}
		x += xStep
	}
}

func countOverlappingInDiagonal(x1 int, y1 int, x2 int, y2 int, floor [][]floorPoint, threshold int, exceededThresholdCount *int) {
	xStep := 1
	if x1 > x2 {
		xStep = -1
	}
	yStep := 1
	if y1 > y2 {
		yStep = -1
	}

	x := x1; y := y1
	breakAfterCheck := false
	for {
		if x == x2 && y == y2 {
			breakAfterCheck = true
		}
		pt := &floor[y][x]
		pt.overlaps += 1
		if !pt.exceededThreshold && pt.overlaps >= threshold {
			pt.exceededThreshold = true
			*exceededThresholdCount++
		}
		x += xStep
		y += yStep

		if breakAfterCheck {
			break
		}
	}
}

func countOverlappingPointsLinesOnly(lines []string, threshold int) int {
	entries, floor := parseLines(lines, true)

	exceededThresholdCount := 0
	for _, e := range entries {
		if e.isDiagonal() {
			continue
		}
		countOverlappingInSingleLine(e.a.x, e.a.y, e.b.x, e.b.y, floor, threshold, &exceededThresholdCount)
	}
	return exceededThresholdCount
}

func countOverlappingPointsLinesAndDiagonals(lines []string, threshold int) int {
	entries, floor := parseLines(lines, false)

	exceededThresholdCount := 0
	for _, e := range entries {
		isDiagonal := e.isDiagonal()
		if isDiagonal {
			countOverlappingInDiagonal(e.a.x, e.a.y, e.b.x, e.b.y, floor, threshold, &exceededThresholdCount)
		} else {
			countOverlappingInSingleLine(e.a.x, e.a.y, e.b.x, e.b.y, floor, threshold, &exceededThresholdCount)
		}
	}
	return exceededThresholdCount
}