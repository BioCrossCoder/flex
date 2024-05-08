package itertools

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	convey.Convey("Call Zip on int slices", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := []int{4, 5, 6}
		arr3 := []int{-1, -2}
		expected := [][]int{
			{1, 4, -1},
			{2, 5, -2},
		}
		result, err := Zip(arr1, arr2, arr3)
		assert.Nil(t, err)
		assert.Equal(t, expected, result.Pour())
		result, err = Zip(arr1)
		assert.Equal(t, common.ErrIllegalParamCount, err)
		assert.Nil(t, result)
	})
}

func TestZipLongest(t *testing.T) {
	convey.Convey("Call ZipLongest on int slices", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := []int{4, 5, 6}
		arr3 := []int{-1, -2}
		expected := [][]int{
			{1, 4, -1},
			{2, 5, -2},
			{3, 6, 0},
		}
		result, err := ZipLongest(arr1, arr2, arr3)
		assert.Nil(t, err)
		assert.Equal(t, expected, result.Pour())
		result, err = ZipLongest(arr1)
		assert.Equal(t, common.ErrIllegalParamCount, err)
		assert.Nil(t, result)
	})
}

func TestZipPair(t *testing.T) {
	convey.Convey("Call ZipPair on 2 int slices", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := []string{"1", "2"}
		expected := []*zipPair[int, string]{
			{1, "1"},
			{2, "2"},
		}
		result := ZipPair(arr1, arr2)
		assert.Equal(t, expected, result.Pour())
	})
}

func TestZipPairLongest(t *testing.T) {
	convey.Convey("Call ZipPairLongest on 2 int slices", t, func() {
		arr1 := []int{1, 2, 3}
		arr2 := []string{"1", "2"}
		expected := []*zipPair[int, string]{
			{1, "1"},
			{2, "2"},
			{3, ""},
		}
		result := ZipPairLongest(arr1, arr2)
		assert.Equal(t, expected, result.Pour())
	})
}
