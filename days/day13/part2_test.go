package day13_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day13"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_Part2(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	exampleOutput, err := test.LoadFixture("example-output")
	if !assert.NoError(t, err) {
		return
	}

	day := day13.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, string(exampleOutput), out+"\n")
}

func TestSolver_Part2_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	exampleOutput, err := test.LoadFixture("real-output")
	if !assert.NoError(t, err) {
		return
	}

	day := day13.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, string(exampleOutput), out+"\n")
}
