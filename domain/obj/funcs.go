package obj

type FunOptional[T any] struct {
	value T
}

func (s FunOptional[T]) GetOrDefault(fun func() T) T {
	if &s.value == nil {
		return fun()
	} else {
		return s.value
	}
}
