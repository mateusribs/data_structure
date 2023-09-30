package datastructures

import (
	"fmt"

	datastructures "github.com/mateusribs/data_structures/linked_list"
)

type Stack struct {
	List datastructures.LinkedList
	Top *datastructures.Node
}

func (s *Stack) Empty() bool {
	return s.List.Empty()
}

func (s *Stack) Push(value int64) {
	s.List.PushFront(value)
	s.Top = s.List.Head
}

func (s *Stack) Pop() (value int64, err error) {
	if !s.Empty(){
		value = s.List.PopFront()
		s.Top = s.List.Head
		return
	}
	
	return value, fmt.Errorf("empty stack")
}