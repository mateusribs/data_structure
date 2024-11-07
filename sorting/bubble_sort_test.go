package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestBubbleSort(t *testing.T) {
	nums := []int{1, 5, 7, 2, 4}

	BubbleSort(nums)

	require.ElementsMatch(t, nums, []int{1, 2, 4, 5, 7})

	nums = []int{30, 10, 1, 103, 204, 29, 11}

	BubbleSort(nums)

	require.ElementsMatch(t, nums, []int{1, 10, 11, 29, 30, 103, 204})
}