package itertools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
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
		assert.Equal(t, common.ErrUnexpectedParamCount, err)
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
		assert.Equal(t, common.ErrUnexpectedParamCount, err)
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

func ExampleZip() {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{-1, -2}
	iter, _ := Zip(arr, arr2, arr3)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 4 -1]
	// [2 5 -2]
}

func ExampleZipLongest() {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{-1, -2}
	iter, _ := ZipLongest(arr, arr2, arr3)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 4 -1]
	// [2 5 -2]
	// [3 6 0]
}

func ExampleZipPair() {
	arr := []int{1, 2, 3}
	arr2 := []string{"a", "b"}
	iter := ZipPair(arr, arr2)
	for iter.Next() {
		fmt.Println(*iter.Value())
	}
	// Output:
	// {1 a}
	// {2 b}
}

func ExampleZipPairLongest() {
	arr := []int{1, 2, 3}
	arr2 := []string{"a", "b"}
	iter := ZipPairLongest(arr, arr2)
	for iter.Next() {
		fmt.Println(*iter.Value())
	}
	// Output:
	// {1 a}
	// {2 b}
	// {3 }
}
