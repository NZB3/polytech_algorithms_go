package doublelinkedlist

import (
	"fmt"
	"reflect"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoubleLinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoubleLinkedList[T any](array ...T) *DoubleLinkedList[T] {
	list := &DoubleLinkedList[T]{nil, nil, 0}
	if len(array) == 0 {
		return list
	}
	list.Fill(array)
	return list
}

func (list *DoubleLinkedList[T]) Fill(array []T) {
	if len(array) == 0 {
		return
	}

	for _, value := range array {
		list.Append(value)
	}
}

func (list *DoubleLinkedList[T]) Len() int {
	return list.length
}

func (list *DoubleLinkedList[T]) Append(value T) {
	node := &Node[T]{value, nil, nil}
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
		node.prev = list.tail
	}
	list.tail = node
	list.length++
}

func (list *DoubleLinkedList[T]) Prepend(value T) {
	node := &Node[T]{value, nil, nil}
	if list.head == nil {
		list.tail = node
	} else {
		list.head.prev = node
		node.next = list.head
	}
	list.head = node
	list.length++
}

func (list *DoubleLinkedList[T]) Remove(index int) (success bool) {
	if index >= list.length {
		return false
	}
	if index == 0 {
		list.head = list.head.next
		list.head.prev = nil
		list.length--
		return true
	}
	current := list.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = current.next.next
	if current.next != nil {
		current.next.prev = current
	}
	list.length--
	return true
}

func (list *DoubleLinkedList[T]) Get(index int) (x T, success bool) {
	if index >= list.length {
		return x, false
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return x, false
}

func (list *DoubleLinkedList[T]) Set(index int, value T) (success bool) {
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

func (list *DoubleLinkedList[T]) Find(value T) (index int) {
	current := list.head
	for i := 0; i < list.length; i++ {
		if reflect.DeepEqual(current.value, value) {
			return i
		}
		current = current.next
	}
	return -1
}

func (list *DoubleLinkedList[T]) Print() {
	current := list.head
	for i := 0; i < list.length; i++ {
		if i == list.length-1 {
			fmt.Print(current.value)
			break
		}
		fmt.Print(current.value, "--")
		current = current.next
	}
	fmt.Println()
}
