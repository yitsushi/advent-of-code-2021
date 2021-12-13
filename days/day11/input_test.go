package day11_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day11"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_SetInput(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day11.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_error(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("input-invalid")
	if !assert.NoError(t, err) {
		return
	}

	day := day11.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	assert.Error(t, err)
}
