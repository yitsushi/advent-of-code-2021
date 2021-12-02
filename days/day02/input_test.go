package day02_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day02"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("example1")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalidCommand(t *testing.T) {
	example, err := test.LoadFixture("input-error")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
}

func TestSolver_SetInput_invalidValue(t *testing.T) {
	example, err := test.LoadFixture("input-error-2")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
}

func TestSolver_SetInput_noValue(t *testing.T) {
	example, err := test.LoadFixture("input-error-3")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
}
