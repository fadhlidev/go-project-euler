package problems_001_to_100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplesOf3And5ShouldReturnCorrectValue(t *testing.T) {
	assert.Equal(t, 543, MultiplesOf3And5(49))
	assert.Equal(t, 233168, MultiplesOf3And5(1000))
	assert.Equal(t, 16687353, MultiplesOf3And5(8456))
	assert.Equal(t, 89301183, MultiplesOf3And5(19564))
}
