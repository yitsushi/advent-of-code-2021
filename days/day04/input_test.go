package day04_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day04"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day04.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_inputError_drawOrder(t *testing.T) {
	example, err := test.LoadFixture("input-error")
	if !assert.NoError(t, err) {
		return
	}

	day := day04.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	assert.Error(t, err)
}

func TestSolver_SetInput_inputError_board(t *testing.T) {
	example, err := test.LoadFixture("input-error-2")
	if !assert.NoError(t, err) {
		return
	}

	day := day04.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	assert.Error(t, err)
}
