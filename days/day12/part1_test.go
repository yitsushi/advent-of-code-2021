package day12_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day12"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_Part1(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "10", out)
}

func TestSolver_Part1_example2(t *testing.T) {
	example, err := test.LoadFixture("example2")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "19", out)
}

func TestSolver_Part1_example3(t *testing.T) {
	example, err := test.LoadFixture("example3")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "226", out)
}

func TestSolver_Part1_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "3410", out)
}
