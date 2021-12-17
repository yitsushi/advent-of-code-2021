package day14_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day14"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_Part2(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "", out)
}

func TestSolver_Part2_real(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "", out)
}
