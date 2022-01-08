package adventofcode21

import (
	"strconv"
	"strings"
	"testing"
)

var sampleInstructions = []instruction{
	{Forward, 5},
	{Down, 5},
	{Forward, 8},
	{Up, 3},
	{Down, 8},
	{Forward, 2},
}

func getFileInstructions() []instruction {
	instructionLines := readStringFile("day2_input")
	instructions := make([]instruction, len(instructionLines))
	for i := 0; i < len(instructionLines); i++ {
		line := instructionLines[i]
		parts := strings.Split(line, " ")
		dir := direction(parts[0])
		step, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{
			Direction: dir,
			Steps:     step,
		}
	}
	return instructions
}

func Test_calculateHorizontalPositionAndDepthInSample(t *testing.T) {
	horPos, depth := calculateHorizontalPositionAndDepth(sampleInstructions)
	act := horPos * depth

	exp := 150
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calculateHorizontalPositionAndDepthInInput(t *testing.T) {
	instructions := getFileInstructions()

	horPos, depth := calculateHorizontalPositionAndDepth(instructions)
	act := horPos * depth

	exp := 1815044
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calculateHorizontalPositionAndDepthWithAimInSample(t *testing.T) {
	horPos, depth := calculateHorizontalPositionAndDepthWithAim(sampleInstructions)
	act := horPos * depth

	exp := 900
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calculateHorizontalPositionAndDepthWithAimInInput(t *testing.T) {
	instructions := getFileInstructions()

	horPos, depth := calculateHorizontalPositionAndDepthWithAim(instructions)
	act := horPos * depth

	exp := 1739283308
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
