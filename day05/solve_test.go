package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	part1, part2 := SolveDay05(input)

	assert.Equal(t, 3, part1)
	assert.Equal(t, 14, part2)
}
