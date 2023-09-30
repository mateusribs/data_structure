package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T){
	s := Stack{}

	require.True(t, s.Empty())
	require.Empty(t, s.Top)

	s.Push(int64(5))
	require.NotEmpty(t, s.Top)
	require.Equal(t, int64(5), s.Top.Value)

	s.Push(int64(4))
	s.Push(int64(7))
	require.Equal(t, int64(7), s.Top.Value)

	value, err := s.Pop()
	require.Equal(t, int64(7), value)
	require.NoError(t, err)

	value, err = s.Pop()
	require.Equal(t, int64(4), value)
	require.NoError(t, err)

	value, err = s.Pop()
	require.NoError(t, err)
	require.Equal(t, int64(5), value)
	require.True(t, s.Empty())
	require.Empty(t, s.Top)

	value, err = s.Pop()
	require.Empty(t, value)
	require.Error(t, err)
}