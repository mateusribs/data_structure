package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestBinarySearch(t *testing.T) {

	array := []int{1, 4, 6, 8, 10, 22, 42}

	index := BinarySearch(array, 6)

	require.Equal(t, index, 2)

	index = BinarySearch(array, 1)

	require.Equal(t, index, 0)

	index = BinarySearch(array, 100)

	require.Equal(t, index, -1)
}