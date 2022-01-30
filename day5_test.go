package adventofcode21

import "testing"

var sampleLines []string = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}


func Test_countOverlappingPointsLinesOnlyInSample(t *testing.T) {

	act := countOverlappingPointsLinesOnly(sampleLines, 2)

	exp := 5
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countOverlappingPointsLinesOnlyInInput(t *testing.T) {
	lines := readStringFile("day5_input")

	act := countOverlappingPointsLinesOnly(lines, 2)

	exp := 4745
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countOverlappingPointsLinesAndDiagonalsInSample(t *testing.T) {
	act := countOverlappingPointsLinesAndDiagonals(sampleLines, 2)

	exp := 12
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countOverlappingPointsLinesAndDiagonalsInInput(t *testing.T) {
	lines := readStringFile("day5_input")

	act := countOverlappingPointsLinesAndDiagonals(lines, 2)

	exp := 18442
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_isDiagonal(t *testing.T)  {
	tests := []struct{
		entryStr string
		exp bool
	} {
		{
			"1,1 -> 3,3", true,
		},
		{
			"9,7 -> 7,9", true,
		},
		{
			"0,9 -> 5,9", false,
		},
		{
			"8,0 -> 0,8", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.entryStr, func(t *testing.T) {
			a, b := getPointsFromLine(tt.entryStr)
			e := entry{
				a: a,
				b: b,
			}
			act := e.isDiagonal()
			if tt.exp != act {
				t.Errorf("expected %t but actual is %t", tt.exp, act)
			}

		})
	}
}