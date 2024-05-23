package list

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIndex(t *testing.T) {
	// Happy path
	assert.Equal(t, SliceIndex(2, 5, true), 2)
	// Edge cases
	assert.Equal(t, SliceIndex(-1, 5, true), 4)
	assert.Equal(t, SliceIndex(5, 5, false), 5)
	assert.Equal(t, SliceIndex(5, 5, true), 4)
	assert.Equal(t, SliceIndex(-7, 10, false), 3)
	assert.Equal(t, SliceIndex(-17, 10, false), -1)
	assert.Equal(t, SliceIndex(-17, 10, true), 0)
}

func TestParseIndex(t *testing.T) {
	// Test happy path
	assert.Equal(t, ParseIndex(2, 5), 2)
	// Test when index is negative
	assert.Equal(t, ParseIndex(-3, 5), 2)
	// Test when index is negative and exceeds length
	assert.Equal(t, ParseIndex(-7, 5), 0)
	// Test when index exceeds length
	assert.Equal(t, ParseIndex(7, 5), 5)
	// Test when index is equal to length
	assert.Equal(t, ParseIndex(5, 5), 5)
}

func TestIsIndexValid(t *testing.T) {
	assert.Nil(t, IsIndexValid(2, 5))
	assert.Equal(t, IsIndexValid(3, 3), common.ErrOutOfRange)
	assert.Equal(t, IsIndexValid(-1, 5), common.ErrOutOfRange)
	assert.Equal(t, IsIndexValid(5, 4), common.ErrOutOfRange)
}
