package adventofcode21

type direction string

const (
	Forward direction = "forward"
	Up      direction = "up"
	Down    direction = "down"
)

type instruction struct {
	Direction direction
	Steps     int
}

func calculateHorizontalPositionAndDepth(instructions []instruction) (int, int) {
	horPos := 0
	depth := 0
	for _, instr := range instructions {
		if instr.Direction == Forward {
			horPos += instr.Steps
		}
		if instr.Direction == Up {
			depth -= instr.Steps
		}
		if instr.Direction == Down {
			depth += instr.Steps
		}
	}
	return horPos, depth
}

func calculateHorizontalPositionAndDepthWithAim(instructions []instruction) (int, int) {
	aim := 0
	horPos := 0
	depth := 0
	for _, instr := range instructions {
		if instr.Direction == Forward {
			horPos += instr.Steps
			depth += aim * instr.Steps
		}
		if instr.Direction == Up {
			aim -= instr.Steps
		}
		if instr.Direction == Down {
			aim += instr.Steps
		}
	}
	return horPos, depth
}
