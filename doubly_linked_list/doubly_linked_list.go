package datastructures

import "errors"

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int64
	Next *Node
	Prev *Node
}

func (l *DoublyLinkedList) Empty() bool {
	return l.Head == nil
}

func (l *DoublyLinkedList) hasOnlySingleNode() bool {
	return l.Head == l.Tail
} 

func (l *DoublyLinkedList) PushFront(value int64) {
	node := Node{Value: value}

	if l.Head != nil {
		node.Next = l.Head
		l.Head.Prev = &node
	}

	l.Head = &node

	if l.Tail == nil {
		l.Tail = l.Head
	}
}

func (l *DoublyLinkedList) PopFront() (value int64) {
	value = l.Head.Value

	if l.Head.Next != nil {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	} else {
		l.Head = nil
		l.Tail = nil
	}

	return
}

func (l *DoublyLinkedList) PushBack(value int64) {
	node := Node{Value: value}

	if l.Empty() {
		l.Head = &node
		l.Tail = &node
	} else {
		l.Tail.Next = &node
		node.Prev = l.Tail
		l.Tail = &node
	}
}

func (l *DoublyLinkedList) PopBack() (value int64) {
	value = l.Tail.Value

	if !l.Empty() {
		if l.hasOnlySingleNode() {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Tail = l.Tail.Prev
			l.Tail.Next = nil
		}
	}
	return
}

func (l *DoublyLinkedList) Find(value int64) (node *Node) {
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

func (l *DoublyLinkedList) Delete(value int64) error {
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
					node.Next.Prev = prev 
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