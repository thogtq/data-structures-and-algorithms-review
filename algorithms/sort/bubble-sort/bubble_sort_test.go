package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	tests := []struct {
		inputArr []int
		expected []int
	}{
		{
			[]int{7, 5, 8, 1, 4, 2, 6, 3},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{1, 2, 3, 9, 5, 77, 12, 4},
			[]int{1, 2, 3, 4, 5, 9, 12, 77},
		},
	}
	for _, test := range tests {
		actual := BubbleSort(test.inputArr)
		assert.Equal(test.expected, actual)
	}
}
