package stack

type Stack[T any] struct {
	keys []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) Push(key T) {
	stack.keys = append(stack.keys, key)
}

func (stack *Stack[T]) Peek() (x T, empty bool) {
	if len(stack.keys) > 0 {
		x = stack.keys[len(stack.keys)-1]
		return x, false
	}
	return x, false
}

func (stack *Stack[T]) Pop() (x T, empty bool) {
	if len(stack.keys) > 0 {
		x = stack.keys[len(stack.keys)-1]
		stack.keys = stack.keys[:len(stack.keys)-1]
		return x, false
	}
	return x, true
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.keys) == 0
}
