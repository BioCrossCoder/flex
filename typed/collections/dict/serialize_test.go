package dict

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	convey.Convey("jsonify and stringify", t, func() {
		d := Dict[string, int]{"a": 1, "b": 2, "c": 3}
		data, err := json.Marshal(d)
		assert.Nil(t, err)
		d2 := Dict[string, int]{}
		assert.Nil(t, json.Unmarshal(data, &d2))
		assert.Equal(t, fmt.Sprint(d), fmt.Sprint(d2))
	})
}
