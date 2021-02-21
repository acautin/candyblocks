package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	// Grid should reserve have defined W and H and be empty
	w := 10
	h := 20
	grid := NewGrid(w, h)
	if assert.NotNil(t, grid, "New grid should not be nil") {
		assert.False(t, grid.HasActiveCandy(), "New grid should be empty")
		assert.Equal(t, w, grid.Width, "New grid should have defined width")
		assert.Equal(t, h, grid.Height, "New grid should have defined height")
	}
}

func TestAddCandy(t *testing.T) {
	generator, grid := initDefaults()
	err := grid.AddCandy(generator.Generate())
	assert.Nil(t, err, "Add first candy should not fail")
	assert.True(t, grid.HasActiveCandy(), "Add candy should set it to active")
	err = grid.AddCandy(generator.Generate())
	assert.NotNil(t, err, "Add candy only possible if grid doesn't have an active candy")
	assert.Equal(t, ErrGridAlreadyHasActiveCandy, err)
}

func initDefaults() (*CandyGenerator, *Grid) {
	return NewCandyGenerator(1), NewGrid(DefaultWidth, DefaultHeight)
}
