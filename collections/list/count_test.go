package list

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCount(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		counts   []int
		expected int
	}{
		{
			name:     "Happy path",
			length:   10,
			counts:   []int{5},
			expected: 5,
		},
		{
			name:     "No counts provided",
			length:   10,
			expected: 1,
		},
		{
			name:     "Zero count provided",
			length:   10,
			counts:   []int{0},
			expected: 10,
		},
	}
	for _, tt := range tests {
		convey.Convey(tt.name, t, func() {
			result := ParseCount(tt.length, tt.counts...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchCount(t *testing.T) {
	// Testing the happy path
	assert.Equal(t, 5, SearchCount(10, 5))
	// Testing the edge cases
	assert.Equal(t, 10, SearchCount(10))
	assert.Equal(t, 10, SearchCount(10, 0))
	assert.Equal(t, 10, SearchCount(10, -1))
}
