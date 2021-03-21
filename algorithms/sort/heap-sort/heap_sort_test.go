package heapsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaptify(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		array    []int
		offset   int
		heapSize int
		expected []int
	}{
		{
			[]int{7, 5, 8, 1, 4, 2, 6, 3},
			3,
			8,
			[]int{7, 5, 8, 3, 4, 2, 6, 1},
		},
		{
			[]int{7, 5, 8, 3, 4, 2, 6, 1},
			0,
			8,
			[]int{8, 5, 7, 3, 4, 2, 6, 1},
		},
		{
			[]int{3, 4, 1, 2, 6, 5},
			2,
			6,
			[]int{3, 4, 5, 2, 6, 1},
		},
		{
			[]int{3, 4, 1, 2, 5, 6},
			0,
			4,
			[]int{4, 3, 1, 2, 5, 6},
		},
	}
	for _, test := range tests {
		heaptify(test.array, test.offset, test.heapSize)
		assert.Equal(test.expected, test.array)
	}
}
func TestBuildHeap(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		array    []int
		expected []int
	}{
		{
			[]int{3, 4, 1, 2, 6, 5},
			[]int{6, 4, 5, 2, 3, 1},
		},
		{
			[]int{7, 5, 8, 1, 4, 2, 6, 3},
			[]int{8, 5, 7, 3, 4, 2, 6, 1},
		},
	}
	for _, test := range tests {
		buildHeap(test.array)
		assert.Equal(test.expected, test.array)
	}
}
func TestHeapSort(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		array    []int
		expected []int
	}{
		{
			[]int{3, 4, 1, 2, 6, 5},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{7, 5, 8, 1, 4, 2, 6, 3},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		HeapSort(test.array)
		assert.Equal(test.expected, test.array)
	}
}
