package game

// DefaultWidth is the default grid width (Number of columns)
const DefaultWidth = 15

// DefaultHeight is the default grid height (Number of rows)
const DefaultHeight = 34

// Grid represents the board where the candy is placed
type Grid struct {
	Width       int
	Height      int
	ActiveCandy *Candy
}

// NewGrid instantiates a new grid with the given width and height
func NewGrid(w int, h int) *Grid {
	return &Grid{
		Width:  w,
		Height: h,
	}
}

// HasActiveCandy checks if the grid has an active candy
func (g *Grid) HasActiveCandy() bool {
	return g.ActiveCandy != nil
}

// AddCandy adds an active candy to the grid, fails if there one already
func (g *Grid) AddCandy(candy *Candy) error {
	if g.HasActiveCandy() {
		return ErrGridAlreadyHasActiveCandy
	}
	g.ActiveCandy = candy
	return nil
}
