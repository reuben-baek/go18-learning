package generic

type Option[T any] interface {
	Get() T
	GetOrElse(v T) T
	IsNone() bool
	ForEach(f func(v T))
}

type List[T any] interface {
	Head() T
	Tail() List[T]
	Len() int
	Prepend(element T) List[T]
}
