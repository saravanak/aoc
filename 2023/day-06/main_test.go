package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitRanges(t *testing.T) {

	var (
		source      = AddressRange{start: 0, end: 100}
		destination = AddressRange{start: 40, end: 45}
	)

	var splitRanges = splitRange(source, destination)

	assert.Equal(t, len(splitRanges), 3, "splits (c)")
	assert.Equal(t, splitRanges[0], AddressRange{start: 0, end: 39})
	assert.Equal(t, splitRanges[1], AddressRange{start: 40, end: 44})
	assert.Equal(t, splitRanges[2], AddressRange{start: 45, end: 100})

	source = AddressRange{start: 40, end: 45}
	destination = AddressRange{start: 0, end: 100}

	splitRanges = splitRange(source, destination)

	assert.Equal(t, len(splitRanges), 1, "splits c()")
	assert.Equal(t, splitRanges[0], AddressRange{start: 40, end: 45})
}

func TestSplitRangesDisjoint(t *testing.T) {

	var (
		source      = AddressRange{start: 0, end: 100}
		destination = AddressRange{start: 101, end: 200}
	)

	var splitRanges = splitRange(source, destination)

	assert.Equal(t, make([]AddressRange, 0), splitRanges)

	source = AddressRange{start: 300, end: 400}
	destination = AddressRange{start: 101, end: 200}

	splitRanges = splitRange(source, destination)

	assert.Equal(t, make([]AddressRange, 0), splitRanges)

}

func TestSplitRangesEndIntersecting(t *testing.T) {

	var (
		source      = AddressRange{start: 0, end: 50}
		destination = AddressRange{start: 45, end: 300}
	)

	var splitRanges = splitRange(source, destination)

	assert.Equal(t, len(splitRanges), 2, "splits (c)")
	assert.Equal(t, splitRanges[0], AddressRange{start: 0, end: 44})
	assert.Equal(t, splitRanges[1], AddressRange{start: 45, end: 50})

}

func TestSplitRangesSourceIntersecting(t *testing.T) {

	var (
		source      = AddressRange{start: 0, end: 50}
		destination = AddressRange{start: 3, end: 400}
	)

	var splitRanges = splitRange(source, destination)

	assert.Equal(t, len(splitRanges), 2, "splits (c)")
	assert.Equal(t, splitRanges[0], AddressRange{start: 0, end: 2})
	assert.Equal(t, splitRanges[1], AddressRange{start: 3, end: 50})

}

func TestSplitBothAsSameRange(t *testing.T) {

	var (
		source      = AddressRange{start: 0, end: 50}
		destination = AddressRange{start: 0, end: 50}
	)

	var splitRanges = splitRange(source, destination)

	assert.Equal(t, len(splitRanges), 1)
	assert.Equal(t, splitRanges[0], AddressRange{start: 0, end: 50})

}
