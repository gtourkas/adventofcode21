package adventofcode21

import "testing"

func Test_calcPowerConsumptionInSample(t *testing.T) {
	measurements := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	act := calcPowerConsumption(measurements)

	exp := int64(198)
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcPowerConsumptionInFile(t *testing.T) {
	measurements := readStringFile("day3_input")

	act := calcPowerConsumption(measurements)

	exp := int64(3374136)
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcLifeSupportRatingInSample(t *testing.T) {
	measurements := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	act := calcLifeSupportRating(measurements)

	exp := int64(230)
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcLifeSupportRatingInFile(t *testing.T) {
	measurements := readStringFile("day3_input")

	act := calcLifeSupportRating(measurements)

	exp := int64(4432698)
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
