package day03_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2021/days/day03"
	test "github.com/yitsushi/go-aoc/testing"
)

func TestSolver_SetInput(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day03.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}
