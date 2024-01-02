package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGradientUp(t *testing.T) {
	var high int
	high = gradientUp(7, 9)
	assert.Equal(t, 5, high)

	high = gradientUp(15, 40)
	assert.Equal(t, 11, high)

	high = gradientUp(30, 200)
	assert.Equal(t, 19, high)

}

func TestGradientDown(t *testing.T) {
	var low int
	low = gradientDown(7, 9)
	assert.Equal(t, 2, low)

	low = gradientDown(15, 40)
	assert.Equal(t, 4, low)

	low = gradientDown(30, 200)
	assert.Equal(t, 11, low)

}
