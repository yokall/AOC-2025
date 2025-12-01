package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveExample(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	part1, part2 := Solve(input)

	assert.Equal(t, 3, part1, "part1 should match expected value")
	assert.Equal(t, 0, part2, "part2 should match expected value")
}
