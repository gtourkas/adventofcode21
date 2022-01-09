package adventofcode21

import "strconv"

func calcPowerConsumption(measurements []string) int64 {

	l := len(measurements[0])

	zeros := 0
	ones := 0
	gammaRateBits := make([]byte, l)
	epsilonRateBits := make([]byte, l)
	for i := 0; i < l; i++ {
		for _, meas := range measurements {
			bit := meas[i]
			if bit == '0' {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			gammaRateBits[i] = '0'
			epsilonRateBits[i] = '1'
		} else {
			gammaRateBits[i] = '1'
			epsilonRateBits[i] = '0'
		}
		zeros = 0
		ones = 0
	}

	gammaRateBitsStr := string(gammaRateBits)
	epsilonRateBitsStr := string(epsilonRateBits)

	gammaRate, _ := strconv.ParseInt(gammaRateBitsStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBitsStr, 2, 64)

	return gammaRate * epsilonRate
}

func findRatingMeasurement(measurements []string, useMostCommonBit bool) string {
	l := len(measurements[0])

	var measurementsWithZero []string
	var measurementsWithOne []string
	zeros := 0
	ones := 0
	for i := 0; i < l; i++ {
		for _, meas := range measurements {
			bit := meas[i]
			if bit == '0' {
				zeros++
				measurementsWithZero = append(measurementsWithZero, meas)
			} else {
				ones++
				measurementsWithOne = append(measurementsWithOne, meas)
			}
		}
		if useMostCommonBit {
			if ones >= zeros {
				measurements = measurementsWithOne
			} else {
				measurements = measurementsWithZero
			}
		} else {
			if ones >= zeros {
				measurements = measurementsWithZero
			} else {
				measurements = measurementsWithOne
			}
		}
		measurementsWithZero = []string{}
		measurementsWithOne = []string{}
		zeros = 0
		ones = 0
		if len(measurements) == 1 {
			break
		}
	}

	return measurements[0]
}

func calcLifeSupportRating(measurements []string) int64 {

	oxygenGeneratorRatingStr := findRatingMeasurement(measurements, true)
	co2ScrubberRatingStr := findRatingMeasurement(measurements, false)

	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenGeneratorRatingStr, 2, 64)
	co2ScrubberRating, _ := strconv.ParseInt(co2ScrubberRatingStr, 2, 64)

	return oxygenGeneratorRating * co2ScrubberRating
}
