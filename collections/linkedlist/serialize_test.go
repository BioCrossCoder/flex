package linkedlist

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	convey.Convey("jsonify and stringify", t, func() {
		l := NewLinkedList(1, 2, 3)
		data, err := json.Marshal(l)
		assert.Nil(t, err)
		l2 := NewLinkedList()
		assert.Nil(t, json.Unmarshal(data, l2))
		assert.Equal(t, fmt.Sprint(l), fmt.Sprint(l2))
	})
}
