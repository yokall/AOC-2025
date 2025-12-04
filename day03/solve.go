package day03

import (
	"strconv"
	"strings"

	"github.com/yokall/aoc-2025/utils"
)

func SolveDay03(input string) (int, int) {
	lines := utils.ReadLinesString(input)

	part1 := solvePart1(lines)
	part2 := 0
	return part1, part2
}

func solvePart1(lines []string) int {
	total := 0
	for _, line := range lines {
		best := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				voltage, _ := strconv.Atoi(strings.Join([]string{string(line[i]), string(line[j])}, ""))
				if voltage > best {
					best = voltage
				}
			}
		}

		total += best
	}

	return total
}
