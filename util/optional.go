package util

type Optional[T any] struct {
	value   T
	present bool
}

func None[T any]() Optional[T] {
	return Optional[T]{
		present: false,
	}
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value:   value,
		present: true,
	}
}

func (o *Optional[T]) IsPresent() bool {
	return o.present
}

func (o *Optional[T]) Get() T {
	return o.value
}
