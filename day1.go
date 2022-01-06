package adventofcode21

func countIncreases(measurements []int) int {
	r := 0
	prev := 0
	for i := 0; i < len(measurements); i++ {
		if i == 0 {
			continue
		}
		cur := measurements[i]
		if cur > prev {
			r++
		}
		prev = cur
	}
	return r
}

func countIncreasesInWindow(measurements []int, windowSize int) int {
	r := 0
	prevWindowSum := 0
	curWindowSum := 0
	for i := 0; i < len(measurements); i++ {
		curWindowSum += measurements[i]
		if i >= windowSize {
			curWindowSum -= measurements[i-windowSize]

			if curWindowSum > prevWindowSum {
				r++
			}
		}
		prevWindowSum = curWindowSum
	}
	return r
}
