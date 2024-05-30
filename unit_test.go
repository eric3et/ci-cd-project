package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageComponentPass(t *testing.T) {
	// 12
	input := []float64{3.14, 3.14, 3.14}

	assert.Equal(t, 3.14, average(input))
}

// func TestAverageComponentFail(t *testing.T) {
// 	input := []float64{3.14, 3.14, 3.14}

// 	assert.Equal(t, 1, average(input))
// }
