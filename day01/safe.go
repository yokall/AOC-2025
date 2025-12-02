package day01

import (
	"math"
	"strconv"
)

type Safe struct {
	maxNumber int
	positions []int
}

// NewSafe creates a new Safe object initialized with position 50
func NewSafe() *Safe {
	return &Safe{
		maxNumber: 99,
		positions: []int{50},
	}
}

func (safe *Safe) TurnDial(instruction string) {
	direction := string(instruction[0])
	distance, _ := strconv.Atoi(instruction[1:])
	distance = distance % (safe.maxNumber + 1)

	switch direction {
	case "L":
		newPosition := safe.positions[len(safe.positions)-1] - distance
		if newPosition < 0 {
			newPosition += safe.maxNumber + 1
		}
		safe.positions = append(safe.positions, newPosition)
	case "R":
		newPosition := safe.positions[len(safe.positions)-1] + distance
		if newPosition > safe.maxNumber {
			newPosition -= safe.maxNumber + 1
		}
		safe.positions = append(safe.positions, newPosition)
	}
}

func (safe *Safe) HowManyTimesZero(instruction string) int {
	count := 0

	direction := string(instruction[0])
	distance, _ := strconv.Atoi(instruction[1:])
	currentPosition := safe.positions[len(safe.positions)-1]

	count += int(math.Abs(float64(distance / 100)))

	distance = distance % 100

	switch direction {
	case "L":
		if currentPosition != 0 && distance > currentPosition {
			count++
		}
	case "R":
		if (currentPosition + distance) > 100 {
			count++
		}
	}

	return count
}
