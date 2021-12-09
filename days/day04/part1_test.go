package day04_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day04"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_Part1(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day04.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "4512", out)
}

func TestSolver_Part1_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day04.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "8442", out)
}
