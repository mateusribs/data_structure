package datastructures

import (
	"errors"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int64
	Next *Node
}

func (l *LinkedList) Empty() bool {
	return l.Head == nil 
}


// search for last node before tail
func (l *LinkedList) beforeTail() (node *Node) {
	node = l.Head
	for node.Next.Next != nil {
		node = node.Next
	}
	return
}

func (l *LinkedList) hasOnlySingleNode() bool {
	return l.Head == l.Tail
}

func (l *LinkedList) PushFront(value int64) {
	node := Node{Value: value}

	if l.Head != nil {
		node.Next = l.Head
	}

	l.Head = &node

	if l.Tail == nil {
		l.Tail = l.Head
	}
}

func (l *LinkedList) PopFront() (value int64) {

	value = l.Head.Value

	if l.Head.Next != nil {
		l.Head = l.Head.Next
	} else {
		l.Head = nil
		l.Tail = nil
	}

	return
}

func (l *LinkedList) PushBack(value int64) {
	
	node := Node{Value: value}

	if l.Empty() {
		l.Head = &node
		l.Tail = &node
	} else {
		l.Tail.Next = &node
		l.Tail = &node
	}
}

func (l *LinkedList) PopBack() (value int64) {

	if !l.Empty() {
		value = l.Tail.Value

		if  l.hasOnlySingleNode() {
			l.Head = nil
			l.Tail = nil
		} else {
			last_node := l.beforeTail()
			last_node.Next = nil
			l.Tail = last_node
		}
	}
	return
}

func (l *LinkedList) Find(value int64) (node *Node) {
	switch value {
	case l.Head.Value:
		return l.Head
	case l.Tail.Value:
		return l.Tail
	default:
		node = l.Head
		for node != nil {
			if node.Value == value {
				return node
			}
			node = node.Next
		}
		return nil
	}
}

func (l *LinkedList) Delete(value int64) error {
	if !l.Empty(){
		switch value {
		case l.Head.Value:
			l.PopFront()
			return nil
		case l.Tail.Value:
			l.PopBack()
			return nil
		default:
			prev := l.Head
			node := l.Head.Next
			for node != nil {
				if node.Value == value {
					prev.Next = node.Next
					return nil
				}
				prev = node
				node = node.Next
			}
			return errors.New("value not found")
		}
	}
	return errors.New("empty list")
}