package itertools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversed(t *testing.T) {
	convey.Convey("reversed array", t, func() {
		arr := [...]int{1, 2, 3, 4, 5}
		r, err := Reversed(arr)
		assert.Nil(t, err)
		assert.Equal(t, []any{5, 4, 3, 2, 1}, r)
	})
	convey.Convey("reversed slice", t, func() {
		s := []any{1, "5", false, nil}
		r, err := Reversed(s)
		assert.Nil(t, err)
		assert.Equal(t, []any{nil, false, "5", 1}, r)
	})
	convey.Convey("reversed string", t, func() {
		s := "hello world"
		r, err := Reversed(s)
		assert.Nil(t, err)
		assert.Equal(t, "dlrow olleh", r)
	})
	convey.Convey("not a sequence", t, func() {
		s := 123
		r, err := Reversed(s)
		assert.Equal(t, err, common.ErrNotSeq)
		assert.Nil(t, r)
	})
}

func ExampleReversed() {
	// string
	s := "hello world"
	r, _ := Reversed(s)
	fmt.Println(r)
	// array
	arr := [...]int{1, 2, 3, 4, 5}
	r, _ = Reversed(arr)
	fmt.Println(r)
	// Output:
	// dlrow olleh
	// [5 4 3 2 1]
}
