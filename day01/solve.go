package day01

import (
	"github.com/yokall/aoc-2025/utils"
)

func Solve(input string) (int, int) {
	lines := utils.ReadLinesString(input)

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	return part1, part2
}

func solvePart1(lines []string) int {
	safe := NewSafe()

	zeroCount := 0
	for _, line := range lines {
		safe.TurnDial(line)

		if safe.positions[len(safe.positions)-1] == 0 {
			zeroCount += 1
		}
	}

	return zeroCount
}

func solvePart2(lines []string) int {
	return 0
}
