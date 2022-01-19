package option

import "go18-learning/generic"

type option[T any] struct {
	value T
	none  bool
}

func (o option[T]) Get() T {
	return o.value
}

func (o option[T]) GetOrElse(v T) T {
	if o.IsNone() {
		return v
	}
	return o.Get()
}

func (o option[T]) IsNone() bool {
	return o.none == true
}

func (o option[T]) ForEach(f func(T)) {
	if o.IsNone() {
		return
	}
	f(o.Get())
}

func Some[T any](v T) generic.Option[T] {
	o := option[T]{
		value: v,
		none:  false,
	}
	return &o
}

func None[T any]() generic.Option[T] {
	o := option[T]{
		none: true,
	}
	return &o
}

func Map[T any, S any](o generic.Option[T], f func(T) S) generic.Option[S] {
	return Some[S](f(o.Get()))
}

func FlatMap[T any, S any](o generic.Option[T], f func(T) generic.Option[S]) generic.Option[S] {
	return Some[S](f(o.Get()).Get())
}
