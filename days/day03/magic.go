package day03

func findMostAndLeastCommonBits(
	list []int64,
	maskLength int,
	mask int64,
) (int64, int64) {
	register := make([]int, maskLength)

	for _, data := range list {
		for idx := range register {
			register[idx] += int(data>>idx) & 1
		}
	}

	var mostCommon int64

	for idx, value := range register {
		if value*2 >= len(list) {
			mostCommon += 1 << idx
		}
	}

	return mostCommon, ^mostCommon & mask
}

func reduceFromGamma(list []int64, maskLength int, mask int64) int64 {
	gammaRate, _ := findMostAndLeastCommonBits(list, maskLength, mask)
	output := []int64{}

	for _, value := range list {
		bit := int64(1 << (maskLength - 1))
		if value&bit == gammaRate&bit {
			output = append(output, value)
		}
	}

	if len(output) == 0 {
		return 0
	}

	if len(output) == 1 {
		return output[0]
	}

	return reduceFromGamma(output, maskLength-1, mask>>1)
}

func reduceFromEpsilon(list []int64, maskLength int, mask int64) int64 {
	_, epsilonRate := findMostAndLeastCommonBits(list, maskLength, mask)
	output := []int64{}

	for _, value := range list {
		bit := int64(1 << (maskLength - 1))
		if value&bit == epsilonRate&bit {
			output = append(output, value)
		}
	}

	if len(output) == 0 {
		return 0
	}

	if len(output) == 1 {
		return output[0]
	}

	return reduceFromEpsilon(output, maskLength-1, mask>>1)
}
