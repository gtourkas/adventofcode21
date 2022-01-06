package adventofcode21

import "testing"

func Test_countIncreasesInSample(t *testing.T) {
	measurements := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	act := countIncreases(measurements)

	exp := 7
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countIncreasesInInput(t *testing.T) {
	measurements := readIntFile("day1_input")

	act := countIncreases(measurements)

	exp := 1832
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countIncreasesInWindowInSample(t *testing.T) {
	measurements := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	act := countIncreasesInWindow(measurements, 3)

	exp := 5
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countIncreasesInWindowInInput(t *testing.T) {
	measurements := readIntFile("day1_input")

	act := countIncreasesInWindow(measurements, 3)

	exp := 1858
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
