package day05

import (
	"strconv"
	"strings"

	"github.com/yokall/aoc-2025/utils"
)

func SolveDay05(input string) (int, int) {
	lines := utils.ReadLinesString(input)
	groups := utils.SplitOn(lines, "")

	freshRanges := groups[0]
	ingredients := groups[1]

	part1 := solvePart1(ingredients, freshRanges)
	part2 := 0

	return part1, part2
}

func solvePart1(ingredients []string, freshRanges []string) int {
	count := 0

	for _, ingredient := range ingredients {
		ingredient, _ := strconv.Atoi(ingredient)
		fresh := false

		for _, freshRange := range freshRanges {
			if inRange(ingredient, freshRange) {
				fresh = true
				break
			}
		}

		if fresh {
			count++
		}
	}

	return count
}

func inRange(ingredient int, freshRange string) bool {
	parts := strings.Split(freshRange, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return ingredient >= start && ingredient <= end
}
