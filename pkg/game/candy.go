package game

import "math/rand"

// CandyType defines the type of candy
//go:generate stringer -type=CandyType
type CandyType uint8

// Enumeration of types of candy based on their shape.
const (
	CandyTypeO CandyType = iota
	CandyTypeI CandyType = iota
	CandyTypeT CandyType = iota
	CandyTypeJ CandyType = iota
	CandyTypeL CandyType = iota
	CandyTypeS CandyType = iota
	CandyTypeZ CandyType = iota
)

// Candy is the main entity of the game and appears randomly at the top.
type Candy struct {
	Type CandyType
}

func candyAvailable() []Candy {
	return []Candy{
		{Type: CandyTypeO},
		{Type: CandyTypeI},
		{Type: CandyTypeT},
		{Type: CandyTypeJ},
		{Type: CandyTypeL},
		{Type: CandyTypeS},
		{Type: CandyTypeZ},
	}
}

// CandyGenerator represents a struct containing the seeded number generator.
type CandyGenerator struct {
	rand    *rand.Rand
	options []Candy
}

// NewCandyGenerator returns a new candy generator initialized with the given seed.
func NewCandyGenerator(seed int64) *CandyGenerator {
	source := rand.NewSource(seed)
	return &CandyGenerator{
		rand:    rand.New(source),
		options: candyAvailable(),
	}
}

// Generate will instantiate a new random candy.
func (c *CandyGenerator) Generate() *Candy {
	return &c.options[c.rand.Intn(len(c.options))]
}
