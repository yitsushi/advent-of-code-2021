package day14

import (
	"math"
)

type transformationMap map[string]string

type polymerMap map[string]int

func resolve(polymer polymerMap, transformations transformationMap) polymerMap {
	result := map[string]int{}

	for pair, count := range polymer {
		trans, found := transformations[pair]
		if !found {
			result[pair] = count

			continue
		}

		result[string(pair[0])+trans] += count
		result[trans+string(pair[1])] += count
	}

	return result
}

func countElements(polymer polymerMap) map[byte]int {
	counter := map[byte]int{}

	for pair, count := range polymer {
		counter[pair[0]] += count
	}

	return counter
}

func minMax(elements map[byte]int) (int, byte, int, byte) {
	var (
		min, max     int = math.MaxInt, 0
		minCh, maxCh byte
	)

	for ch, count := range elements {
		if count > max {
			max = count
			maxCh = ch
		}

		if count < min {
			min = count
			minCh = ch
		}
	}

	return min, minCh, max, maxCh
}
