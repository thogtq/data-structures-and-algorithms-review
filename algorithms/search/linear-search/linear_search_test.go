package linearsearch

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	assert := assert.New(t)
	testArr := []int{3, 4, 2, 9, 8, 11, 7, 23, 6}
	testNumberArr := []int{3, 8, 13, 22, 17, 6, 0, -1, math.MaxInt64}
	testExpectedArr := []int{0, 4, -1, -1, -1, 8, -1, -1, -1}
	for idx, testNumber := range testNumberArr {
		result := LinearSearch(testArr, testNumber)
		assert.Equal(testExpectedArr[idx], result)
	}
}
