package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		inputArr    []int
		inputNumber int
		expected    int
	}{
		{
			[]int{2, 3, 4, 6, 7, 8, 9, 11, 23},
			2,
			0,
		},
		{
			[]int{2, 3, 4, 6, 7, 8, 9, 11, 23},
			23,
			8,
		},
		{
			[]int{2, 3, 4, 6, 7, 8, 9, 11, 23},
			7,
			4,
		},
		{
			[]int{2, 3, 4, 6, 7, 8, 9, 11, 23},
			5,
			-1,
		},
		{
			[]int{1, 7, 19, 33, 79, 126},
			1,
			0,
		},
		{
			[]int{1, 7, 19, 33, 79, 126},
			126,
			5,
		},
		{
			[]int{1, 7, 19, 33, 79, 126},
			33,
			3,
		},
		{
			[]int{1, 7, 19, 33, 79, 126},
			20,
			-1,
		},
	}
	for _, test := range tests {
		actual := BinarySearch(test.inputArr, test.inputNumber)
		assert.Equal(test.expected, actual)
	}
}
