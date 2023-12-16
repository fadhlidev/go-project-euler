package problems_001_to_100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiboEvenSumShouldReturnCorrectValue(t *testing.T) {
	assert.Equal(t, 10, FiboEvenSum(8))
	assert.Equal(t, 10, FiboEvenSum(10))
	assert.Equal(t, 44, FiboEvenSum(34))
	assert.Equal(t, 44, FiboEvenSum(60))
	assert.Equal(t, 798, FiboEvenSum(1000))
	assert.Equal(t, 60696, FiboEvenSum(100000))
	assert.Equal(t, 4613732, FiboEvenSum(4000000))
}
