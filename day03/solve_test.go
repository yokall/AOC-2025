package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	part1, part2 := SolveDay03(input)
	assert.Equal(t, 357, part1)
	assert.Equal(t, 0, part2)
}
