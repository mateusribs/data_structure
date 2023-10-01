package datastructures

import (
	"fmt"

	datastructures "github.com/mateusribs/data_structures/linked_list"
)


type Queue struct {
	List datastructures.LinkedList
}

func (q *Queue) Empty() bool {
	return q.List.Empty()
}

func (q *Queue) Enqueue(value int64) {
	q.List.PushBack(value)
}

func (q *Queue) Dequeue() (value int64, err error) {
	if !q.Empty(){
		value = q.List.PopFront()
		return
	}

	return value, fmt.Errorf("empty queue")
}