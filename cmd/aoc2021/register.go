package main

import (
	"github.com/yitsushi/go-aoc/puzzle"

	"github.com/yitsushi/advent-of-code-2021/days/day01"
	"github.com/yitsushi/advent-of-code-2021/days/day02"
)

func withDays(month puzzle.Month) puzzle.Month {
	month.Register(1, &day01.Solver{})
	month.Register(2, &day02.Solver{})

	return month
}
