package main

import (
	"github.com/yitsushi/go-aoc/puzzle"

	"github.com/yitsushi/advent-of-code-2021/days/day01"
)

func withDays(month puzzle.Month) puzzle.Month {
	month.Register(1, &day01.Solver{})

	return month
}
