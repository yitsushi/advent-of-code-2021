package day02_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day02"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_Part2(t *testing.T) {
	example, err := test.LoadFixture("example1")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "900", out)
}

func TestSolver_Part2_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "1975421260", out)
}

func TestSolver_Part2_inputError(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("input-error")
	if !assert.NoError(t, err) {
		return
	}

	day := day02.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	assert.Error(t, err)
}
