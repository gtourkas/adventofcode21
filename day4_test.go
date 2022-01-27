package adventofcode21

import "testing"

func Test_calcBingoScoreInSample_winnerIsTheFirst(t *testing.T) {
	lines := readStringFile("day4_input_sample")

	act := calcBingoScore(lines, false)

	exp := 4512
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcBingoScoreInFile_winnerIsTheFirst(t *testing.T) {
	lines := readStringFile("day4_input")

	act := calcBingoScore(lines, false)

	exp := 49860
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcBingoScoreInSample_winnerIsTheLast(t *testing.T) {
	lines := readStringFile("day4_input_sample")

	act := calcBingoScore(lines, true)

	exp := 1924
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcBingoScoreInFile_winnerIsTheLast(t *testing.T) {
	lines := readStringFile("day4_input")

	act := calcBingoScore(lines, true)

	exp := 24628
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}