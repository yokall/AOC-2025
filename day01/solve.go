package day01

import (
	"github.com/yokall/aoc-2025/utils"
)

func Solve(input string) (int, int) {
	lines := utils.ReadLinesString(input)
	lines = utils.FilterEmptyLines(lines)

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	return part1, part2
}

func solvePart1(lines []string) int {
	return 0
}

func solvePart2(lines []string) int {
	return 0
}
