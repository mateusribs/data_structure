package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestLinkedList(t *testing.T){
	l := LinkedList{}
	require.Empty(t, l.Head)
	require.Empty(t, l.Tail)

	l.PushFront(int64(64))
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)
	
	l.PushFront(int64(423))
	require.Equal(t, l.Head.Value, int64(423))
	require.Equal(t, l.Tail.Value, int64(64))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	value := l.PopFront()
	require.Equal(t, int64(423), value)
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	value = l.PopFront()
	require.Equal(t, int64(64), value)
	require.Empty(t, l.Head)
	require.Empty(t, l.Tail)

	require.True(t, l.Empty())

	l.PushBack(int64(64))
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	l.PushBack(int64(56))
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Tail.Value, int64(56))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	l.PushBack(int64(5))
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Tail.Value, int64(5))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	value = l.PopBack()
	require.Equal(t, int64(5), value)
	require.Equal(t, l.Head.Value, int64(64))
	require.Equal(t, l.Tail.Value, int64(56))
	require.NotEqual(t, l.Head, l.Tail)
	require.NotEmpty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	value = l.PopBack()
	require.Equal(t, int64(56), value)
	require.Equal(t, l.Head, l.Tail)
	require.Empty(t, l.Head.Next)
	require.Empty(t, l.Tail.Next)

	value = l.PopBack()
	require.Equal(t, int64(64), value)
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
	require.Equal(t, int64(4), l.Head.Next.Next.Value)

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
	require.Equal(t, int64(2), l.Head.Next.Next.Value)

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