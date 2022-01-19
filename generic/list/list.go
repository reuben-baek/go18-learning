package list

import "go18-learning/generic"

type list[T any] struct {
	head T
	tail generic.List[T]
	len  int
}

func (l *list[T]) Len() int {
	return l.len
}
func (l *list[T]) Head() T {
	return l.head
}

func (l *list[T]) Tail() generic.List[T] {
	return l.tail
}

func (l *list[T]) Prepend(element T) generic.List[T] {
	ll := list[T]{
		element,
		l,
		l.len + 1,
	}
	return &ll
}

func (l *list[T]) ForEach(f func(v T)) {
}

func New[T any](items ...T) generic.List[T] {
	var ll generic.List[T]
	l := list[T]{}
	ll = &l
	for i := len(items) - 1; i >= 0; i-- {
		ll = ll.Prepend(items[i])
	}
	return ll
}
