package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveExample(t *testing.T) {
	input := `example line 1
example line 2
example line 3`

	part1, part2 := Solve(input)

	assert.Equal(t, 0, part1, "part1 should match expected value")
	assert.Equal(t, 0, part2, "part2 should match expected value")
}

func TestSolvePart1(t *testing.T) {
	input := []string{"example data"}
	expected := 0

	result := solvePart1(input)

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	input := []string{"example data"}
	expected := 0

	result := solvePart2(input)

	assert.Equal(t, expected, result)
}
