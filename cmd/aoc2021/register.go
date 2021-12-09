package main

import (
	"github.com/yitsushi/go-aoc/puzzle"

	"github.com/yitsushi/advent-of-code-2021/days/day01"
	"github.com/yitsushi/advent-of-code-2021/days/day02"
	"github.com/yitsushi/advent-of-code-2021/days/day03"
	"github.com/yitsushi/advent-of-code-2021/days/day04"
	"github.com/yitsushi/advent-of-code-2021/days/day05"
	"github.com/yitsushi/advent-of-code-2021/days/day06"
)

func withDays(month puzzle.Month) puzzle.Month {
	month.Register(1, &day01.Solver{})
	month.Register(2, &day02.Solver{})
	month.Register(3, &day03.Solver{})
	month.Register(4, &day04.Solver{})
	month.Register(5, &day05.Solver{})
	month.Register(6, &day06.Solver{})

	return month
}
