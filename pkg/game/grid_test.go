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
		assert.True(t, grid.IsEmpty(), "New grid should be empty")
		assert.Equal(t, w, grid.Width, "New grid should have defined width")
		assert.Equal(t, h, grid.Height, "New grid should have defined height")
	}
}
