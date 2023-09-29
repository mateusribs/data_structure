package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoublyLinkedList(t *testing.T){
	l := DoublyLinkedList{}
	require.Empty(t, l.Head)
	require.Empty(t, l.Tail)

	l.PushFront(int64(64))
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)
	require.Empty(t, l.Head.Prev)
	require.Empty(t, l.Tail.Prev)
	
	l.PushFront(int64(423))
	require.Equal(t, l.Head.Value, int64(423))
	require.Equal(t, l.Tail.Value, int64(64))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Head.Prev)
	require.Empty(t, l.Tail.Next)
	require.NotEmpty(t, l.Tail.Prev)

	l.PushFront(int64(55))

	value := l.PopFront()
	require.Equal(t, int64(55), value)
	require.Equal(t, l.Head.Value, int64(423))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Head.Prev)
	require.Empty(t, l.Tail.Next)
	require.NotEmpty(t, l.Tail.Prev)

	l.PopFront()
	value = l.PopFront()
	require.Equal(t, int64(64), value)
	require.Empty(t, l.Head)
	require.Empty(t, l.Tail)
	require.True(t, l.Empty())

	l.PushBack(int64(5))
	require.Equal(t, l.Head.Value, int64(5))
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)
	require.Empty(t, l.Head.Prev)
	require.Empty(t, l.Tail.Prev)

	l.PushBack(int64(7))
	require.Equal(t, l.Head.Value, int64(5))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Equal(t, l.Head.Next.Value, int64(7))
	require.Empty(t, l.Tail.Next)
	require.Empty(t, l.Head.Prev)
	require.NotEmpty(t, l.Tail.Prev)
	require.Equal(t, l.Tail.Prev.Value, int64(5))

	l.PushBack(int64(67))
	require.Equal(t, l.Head.Value, int64(5))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Equal(t, l.Head.Next.Value, int64(7))
	require.Empty(t, l.Tail.Next)
	require.Empty(t, l.Head.Prev)
	require.NotEmpty(t, l.Tail.Prev)
	require.Equal(t, l.Tail.Prev.Value, int64(7))

	value = l.PopBack()
	require.Equal(t, int64(67), value)
	require.Equal(t, l.Tail.Value, int64(7))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Head.Prev)
	require.Empty(t, l.Tail.Next)
	require.NotEmpty(t, l.Tail.Prev)

	l.PopBack()
	value = l.PopBack()
	require.Equal(t, int64(5), value)
	require.Empty(t, l.Head)
	require.Empty(t, l.Tail)
	require.True(t, l.Empty())

	for i := 0; i < 5; i++ {
		l.PushBack(int64(i+1))
	}

	node := l.Find(int64(5))
	require.Equal(t, int64(5), node.Value)


	node = l.Find(int64(1))
	require.Equal(t, int64(1), node.Value)

	node = l.Find(int64(3))
	require.Equal(t, int64(3), node.Value)

	node = l.Find(int64(55))
	require.Empty(t, node)

	err := l.Delete(int64(3))
	require.NoError(t, err)
	require.Equal(t, int64(2), l.Tail.Prev.Prev.Value)

	err = l.Delete(int64(5))
	require.NoError(t, err)
	require.Equal(t, int64(4), l.Tail.Value)

	err = l.Delete(int64(1))
	require.NoError(t, err)
	require.Equal(t, int64(2), l.Head.Value)

	err = l.Delete(int64(54))
	require.Error(t, err)

	l.Delete(int64(2))
	require.Equal(t, int64(4), l.Tail.Value)
	require.Equal(t, int64(4), l.Head.Value)
	l.Delete(int64(4))
	require.True(t, l.Empty())

	err = l.Delete(int64(4))
	require.Error(t, err)

	for i := 0; i < 5; i++ {
		l.PushFront(int64(i+1))
	}

	err = l.Delete(int64(3))
	require.NoError(t, err)
	require.Equal(t, int64(4), l.Tail.Prev.Prev.Value)

	err = l.Delete(int64(5))
	require.NoError(t, err)
	require.Equal(t, int64(4), l.Head.Value)

	err = l.Delete(int64(1))
	require.NoError(t, err)
	require.Equal(t, int64(2), l.Tail.Value)

	err = l.Delete(int64(54))
	require.Error(t, err)

	l.Delete(int64(2))
	require.Equal(t, int64(4), l.Tail.Value)
	require.Equal(t, int64(4), l.Head.Value)
	l.Delete(int64(4))
	require.True(t, l.Empty())

	err = l.Delete(int64(4))
	require.Error(t, err)
}