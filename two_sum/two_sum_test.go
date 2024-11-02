package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestTwoSumHashMap(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	output := TwoSumHashMap(nums, target)

	require.ElementsMatch(t, output, []int{1, 0})

	nums = []int{3, 2, 4}
	target = 6

	output = TwoSumHashMap(nums, target)

	require.ElementsMatch(t, output, []int{2, 1})

	nums = []int{3, 3}
	target = 6

	output = TwoSumHashMap(nums, target)

	require.ElementsMatch(t, output, []int{1, 0})
}


func TestTwoSumBruteForce(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	output := TwoSumBruteForce(nums, target)

	require.ElementsMatch(t, output, []int{1, 0})

	nums = []int{3, 2, 4}
	target = 6

	output = TwoSumBruteForce(nums, target)

	require.ElementsMatch(t, output, []int{2, 1})

	nums = []int{3, 3}
	target = 6

	output = TwoSumBruteForce(nums, target)

	require.ElementsMatch(t, output, []int{1, 0})
}