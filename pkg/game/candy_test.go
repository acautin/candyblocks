package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const seed = 1

func TestCandyGenerator(t *testing.T) {
	sg := NewCandyGenerator(seed)
	assert.NotNil(t, sg, "Candy generator should not be nil")
	expectedTypes := []CandyType{
		CandyTypeZ,
		CandyTypeT,
		CandyTypeI,
	}
	for i := 0; i < len(expectedTypes); i++ {
		candy := sg.Generate()
		assert.NotNil(t, candy, "Generated Candy should not be nil")
		assert.Equalf(t, expectedTypes[i], candy.Type, "Unexpected candyType got:%s expected:%s", candy.Type, expectedTypes[i])
	}

}
