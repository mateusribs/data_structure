package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountNumberOccurrence(t *testing.T) {
	nums := []int{2, 5, 5, 5, 6, 6, 8, 9, 9, 9}
	target := 5

	output := CountNumberOccurence(nums, target)

	require.Equal(t, output, 3)

	target = 6

	output = CountNumberOccurence(nums, target)

	require.Equal(t, output, 2)

	target = 2

	output = CountNumberOccurence(nums, target)

	require.Equal(t, output, 1)

	target = 100

	output = CountNumberOccurence(nums, target)

	require.Equal(t, output, 0)
}