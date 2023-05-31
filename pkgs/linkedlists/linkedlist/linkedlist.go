package linkedlist

import (
	"fmt"
	"reflect"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func NewLinkedList[T any](array ...T) *LinkedList[T] {
	list := &LinkedList[T]{nil, 0}
	if len(array) == 0 {
		return list
	}
	list.Fill(array)
	return list
}

func (list *LinkedList[T]) Fill(array []T) {
	if len(array) == 0 {
		return
	}

	for _, value := range array {
		list.Append(value)
	}
}

func (list *LinkedList[T]) Len() int {
	return list.length
}

func (list *LinkedList[T]) Append(value T) {
	node := &Node[T]{value, nil}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	list.length++
}

func (list *LinkedList[T]) Prepend(value T) {
	node := &Node[T]{value, nil}
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.length++
}

func (list *LinkedList[T]) Remove(index int) (success bool) {
	if index >= list.length {
		return false
	}
	if index == 0 {
		list.head = list.head.next
		list.length--
		return true
	}
	current := list.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	list.length--
	return true
}

func (list *LinkedList[T]) Get(index int) (value T, success bool) {
	if index >= list.length {
		return value, false
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value, true
}

func (list *LinkedList[T]) Set(index int, value T) (success bool) {
	if index >= list.length {
		return false
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = value
	return true
}

func (list *LinkedList[T]) Find(value T) (index int) {
	current := list.head
	for i := 0; i < list.length; i++ {
		if reflect.DeepEqual(current.value, value) {
			return i
		}
		current = current.next
	}
	return -1
}

func (list *LinkedList[T]) Print() {
	current := list.head
	for i := 0; i < list.length; i++ {
		if i == list.length-1 {
			fmt.Print(current.value)
			break
		}
		fmt.Print(current.value, "->")
		current = current.next
	}
	println()
}
