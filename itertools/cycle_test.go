// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycle(t *testing.T) {
	convey.Convey("cycle array", t, func() {
		entry := [3]int{1, 2, 3}
		c, err := Cycle(entry)
		assert.Nil(t, err)
		for i := 0; i < 3; i++ {
			for j := 0; j < len(entry); j++ {
				assert.Equal(t, entry[j], c.Next())
			}
		}
	})
	convey.Convey("cycle slice", t, func() {
		entry := []int{1, 2, 3}
		c, err := Cycle(entry)
		assert.Nil(t, err)
		for i := 0; i < 3; i++ {
			for j := 0; j < len(entry); j++ {
				assert.Equal(t, entry[j], c.Next())
			}
		}
	})
	convey.Convey("cycle string", t, func() {
		entry := "abc"
		c, err := Cycle(entry)
		assert.Nil(t, err)
		for i := 0; i < 3; i++ {
			for j := 0; j < len(entry); j++ {
				assert.Equal(t, string(entry[j]), c.Next())
			}
		}
	})
}
