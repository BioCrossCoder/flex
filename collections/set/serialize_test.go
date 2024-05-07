package set

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestSerialize(t *testing.T) {
	convey.Convey("jsonify and stringify", t, func() {
		l := Of(1, 2, 3)
		data, err := json.Marshal(&l)
		assert.Nil(t, err)
		l2 := Of()
		assert.Nil(t, json.Unmarshal(data, &l2))
		chars1 := strings.Split(fmt.Sprint(l), "")
		sort.Strings(chars1)
		chars2 := strings.Split(fmt.Sprint(l2), "")
		sort.Strings(chars2)
		assert.Equal(t, chars1, chars2)
	})
}
