package day04

import (
	"github.com/yokall/aoc-2025/utils"
)

func SolveDay04(input string) (int, int) {
	grid := utils.Read2DGridFromString(input)

	part1 := solvePart1(grid)
	part2 := solvePart2(grid)
	return part1, part2
}

func solvePart1(lines [][]rune) int {
	total := 0
	for y := range lines {
		for x := 0; x < len(lines[0]); x++ {
			cell := lines[y][x]

			if cell == '@' {

				adjacentRolls := 0
				directions := [][2]int{
					{-1, -1}, // upleft
					{-1, 0},  // up
					{-1, 1},  // upright
					{1, -1},  // downleft
					{1, 0},   // down
					{1, 1},   // downright
					{0, -1},  // left
					{0, 1},   // right
				}

				for _, dir := range directions {
					newY := y + dir[0]
					newX := x + dir[1]

					if newY >= 0 && newY < len(lines) && newX >= 0 && newX < len(lines[0]) {
						if lines[newY][newX] == '@' {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls < 4 {
					total++
				}
			}
		}
	}

	return total
}

func solvePart2(lines [][]rune) int {
	removed := removeRolls(lines, [][2]int{})

	return len(removed)
}

func removeRolls(lines [][]rune, removed [][2]int) [][2]int {
	current := [][2]int{}
	for y := range lines {
		for x := 0; x < len(lines[0]); x++ {
			cell := lines[y][x]

			if cell == '@' {

				adjacentRolls := 0
				directions := [][2]int{
					{-1, -1}, // upleft
					{-1, 0},  // up
					{-1, 1},  // upright
					{1, -1},  // downleft
					{1, 0},   // down
					{1, 1},   // downright
					{0, -1},  // left
					{0, 1},   // right
				}

				for _, dir := range directions {
					newY := y + dir[0]
					newX := x + dir[1]

					if newY >= 0 && newY < len(lines) && newX >= 0 && newX < len(lines[0]) {
						if lines[newY][newX] == '@' {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls < 4 {
					current = append(current, [2]int{y, x})
				}
			}
		}
	}

	if len(current) > 0 {
		for _, coord := range current {
			lines[coord[0]][coord[1]] = '.'
		}

		removed = removeRolls(lines, append(removed, current...))
	}

	return removed
}
