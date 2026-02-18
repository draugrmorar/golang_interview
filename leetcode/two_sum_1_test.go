package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, []int{1, 0}, actual)

	actual = twoSum([]int{3, 2, 4}, 6)
	assert.Equal(t, []int{2, 1}, actual)

	actual = twoSum([]int{3, 3}, 6)
	assert.Equal(t, []int{1, 0}, actual)
}
