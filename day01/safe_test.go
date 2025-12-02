package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSafe(t *testing.T) {
	safe := NewSafe()

	assert.Equal(t, 50, safe.positions[0])
}

func TestTurnDial(t *testing.T) {
	safe := NewSafe()

	// the dial starts at 50
	assert.Equal(t, 50, safe.positions[0])

	// turning it to the left by 10 should result in position 40
	safe.TurnDial("L10")
	assert.Equal(t, 40, safe.positions[len(safe.positions)-1])

	// turning it to the right by 20 should result in position 60
	safe.TurnDial("R20")
	assert.Equal(t, 60, safe.positions[len(safe.positions)-1])

	// turning it to the left by 70 should wrap around and result in position 90
	safe.TurnDial("L70")
	assert.Equal(t, 90, safe.positions[len(safe.positions)-1])

	// turning it to the right by 15 should wrap around and result in position 5
	safe.TurnDial("R15")
	assert.Equal(t, 5, safe.positions[len(safe.positions)-1])

	// turning it to the left by more than 100 should wrap around correctly
	safe.TurnDial("L110")
	assert.Equal(t, 95, safe.positions[len(safe.positions)-1])

	// turning it to the right by more than 100 should wrap around correctly
	safe.TurnDial("R120")
	assert.Equal(t, 15, safe.positions[len(safe.positions)-1])
}

func TestHowManyTimesZero(t *testing.T) {
	safe := NewSafe()

	count := safe.HowManyTimesZero("L68")
	assert.Equal(t, 1, count)

	count = safe.HowManyTimesZero("R68")
	assert.Equal(t, 1, count)

	count = safe.HowManyTimesZero("L168")
	assert.Equal(t, 2, count)

	count = safe.HowManyTimesZero("R168")
	assert.Equal(t, 2, count)

	count = safe.HowManyTimesZero("L268")
	assert.Equal(t, 3, count)

	count = safe.HowManyTimesZero("R268")
	assert.Equal(t, 3, count)
}
