package gowriter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	var array []int
	func(array *[]int) {
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
		*array = append(*array, 1)
	}(&array)
	assert.Len(t, array, 10)
}
