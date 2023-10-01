package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T){
	q := Queue{}
	require.True(t, q.Empty())

	q.Enqueue(int64(3))
	require.False(t, q.Empty())
	require.Equal(t, q.List.Head.Value, int64(3))
	
	q.Enqueue(int64(54))
	q.Enqueue(int64(7))
	require.Equal(t, q.List.Head.Value, int64(3))
	require.Equal(t, q.List.Tail.Value, int64(7))

	value, err := q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, int64(3), value)
	require.Equal(t, int64(54), q.List.Head.Value)

	q.Dequeue()
	value, err = q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, int64(7), value)
	
	value, err = q.Dequeue()
	require.Error(t, err)
	require.Empty(t, value)
}