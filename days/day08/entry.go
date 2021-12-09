// Package day08 monster.
//nolint:gomnd // This file has a shit ton of "magic" numbers but the they are not.
package day08

import (
	"math"

	"github.com/yitsushi/go-aoc/slice"
)

//  0000
// 1    2
// 1    2
//  3333
// 4    5
// 4    5
//  6666

// Entry is one entry in the input.
type Entry struct {
	UniquePatterns     []string
	OutputValue        []string
	NumericOutputValue []int
	Segments           []rune
	Digits             []string
}

// NewEntry creates a new entry with initialised slices.
func NewEntry() Entry {
	return Entry{
		UniquePatterns:     []string{},
		OutputValue:        []string{},
		NumericOutputValue: []int{},
		Segments:           make([]rune, 7),
		Digits:             make([]string, 10),
	}
}

func (entry *Entry) findFixNumbers() {
	for _, digit := range entry.UniquePatterns {
		if entry.IsOne(digit) {
			entry.Digits[1] = digit
		}

		if entry.IsSeven(digit) {
			entry.Digits[7] = digit
		}

		if entry.IsFour(digit) {
			entry.Digits[4] = digit
		}

		if entry.IsEight(digit) {
			entry.Digits[8] = digit
		}
	}
}

func (entry *Entry) findSegment0() {
	rDiff := RuneDiff(entry.Digits[1], entry.Digits[7])
	entry.Segments[0] = rDiff[0]
}

func (entry *Entry) findSegment1() {
	rDiff := RuneDiff(
		entry.Digits[9],
		string([]rune{
			entry.Segments[0],
			entry.Segments[2],
			entry.Segments[3],
			entry.Segments[5],
			entry.Segments[6],
		}),
	)
	entry.Segments[1] = rDiff[0]
}

func (entry *Entry) findSegment2() {
	for _, digit := range entry.UniquePatterns {
		if _, found := slice.ContainsString(entry.Digits, digit); found {
			continue
		}

		rDiff := RuneDiff(digit+entry.Digits[1], entry.Digits[9])
		if len(rDiff) == 0 {
			rDiff = RuneDiff(digit, entry.Digits[9])
			entry.Segments[2] = rDiff[0]
			entry.Digits[5] = digit

			return
		}
	}
}

func (entry *Entry) findSegment3() {
	base := entry.Digits[7] + string(entry.Segments[6])

	for _, digit := range entry.UniquePatterns {
		if len(digit) == 5 {
			rDiff := RuneDiff(base, digit)
			if len(rDiff) == 1 {
				entry.Segments[3] = rDiff[0]
			}
		}
	}
}

func (entry *Entry) findSegment4() {
	rDiff := RuneDiff(entry.Digits[9], entry.Digits[8])
	entry.Segments[4] = rDiff[0]
}

func (entry *Entry) findSegment5() {
	rDiff := RuneDiff(entry.Digits[1], string(entry.Segments[2]))
	entry.Segments[5] = rDiff[0]
}

func (entry *Entry) findSegment6() {
	base := entry.Digits[4] + string(entry.Segments[0])

	for _, digit := range entry.UniquePatterns {
		if len(digit) == 6 {
			rDiff := RuneDiff(base, digit)
			if len(rDiff) == 1 {
				entry.Segments[6] = rDiff[0]
				entry.Digits[9] = digit
			}
		}
	}
}

// TranslateOutput translates output values into digits.
func (entry *Entry) TranslateOutput() {
	for _, value := range entry.OutputValue {
		entry.NumericOutputValue = append(entry.NumericOutputValue, entry.translate(value))
	}
}

func hasWrapper(value string) func(find rune) bool {
	return func(find rune) bool {
		for _, char := range value {
			if char == find {
				return true
			}
		}

		return false
	}
}

// IsZero checks if it's a zero.
func (entry *Entry) IsZero(value string) bool {
	has := hasWrapper(value)

	return len(value) == 6 && !has(entry.Segments[3])
}

// IsOne checks if it's a one.
func (entry *Entry) IsOne(value string) bool {
	return len(value) == 2
}

// IsTwo checks if it's a two.
func (entry *Entry) IsTwo(value string) bool {
	has := hasWrapper(value)

	return len(value) == 5 && !has(entry.Segments[1]) && !has(entry.Segments[5])
}

// IsThree checks if it's a three.
func (entry *Entry) IsThree(value string) bool {
	has := hasWrapper(value)

	return len(value) == 5 && !has(entry.Segments[1]) && !has(entry.Segments[4])
}

// IsFour checks if it's a four.
func (entry *Entry) IsFour(value string) bool {
	return len(value) == 4
}

// IsFive checks if it's a five.
func (entry *Entry) IsFive(value string) bool {
	has := hasWrapper(value)

	return len(value) == 5 && !has(entry.Segments[2]) && !has(entry.Segments[4])
}

// IsSix checks if it's a six.
func (entry *Entry) IsSix(value string) bool {
	has := hasWrapper(value)

	return len(value) == 6 && !has(entry.Segments[2])
}

// IsSeven checks if it's a seven.
func (entry *Entry) IsSeven(value string) bool {
	return len(value) == 3
}

// IsEight checks if it's a eight.
func (entry *Entry) IsEight(value string) bool {
	return len(value) == 7
}

// IsNine checks if it's a nine.
func (entry *Entry) IsNine(value string) bool {
	has := hasWrapper(value)

	return len(value) == 6 && !has(entry.Segments[4])
}

func (entry *Entry) translate(value string) int {
	if entry.IsZero(value) {
		return 0
	}

	if entry.IsOne(value) {
		return 1
	}

	if entry.IsTwo(value) {
		return 2
	}

	if entry.IsThree(value) {
		return 3
	}

	if entry.IsFour(value) {
		return 4
	}

	if entry.IsFive(value) {
		return 5
	}

	if entry.IsSix(value) {
		return 6
	}

	if entry.IsSeven(value) {
		return 7
	}

	if entry.IsEight(value) {
		return 8
	}

	if entry.IsNine(value) {
		return 9
	}

	return -1
}

// NumericValue retruns with the numeric value of the output.
func (entry *Entry) NumericValue() int64 {
	output := int64(0)
	length := len(entry.NumericOutputValue) - 1

	for idx, value := range entry.NumericOutputValue {
		output += int64(value) * int64(math.Pow(10, float64(length-idx)))
	}

	return output
}

// Parse the entry.
func (entry *Entry) Parse() {
	entry.findFixNumbers()
	entry.findSegment0()
	entry.findSegment6()
	entry.findSegment3()
	entry.findSegment2()
	entry.findSegment5()
	entry.findSegment1()
	entry.findSegment4()
}
