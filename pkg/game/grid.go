package game

// Grid represents the board where the candy is placed
type Grid struct {
	Width  int
	Height int
}

// NewGrid instantiates a new grid with the given width and height
func NewGrid(w int, h int) *Grid {
	return &Grid{
		Width:  w,
		Height: h,
	}
}

// IsEmpty checks if the grid is empty
func (g *Grid) IsEmpty() bool {
	return true
}
