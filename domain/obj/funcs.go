package obj

type FunObj[T any] struct {
	value *T
}

func NewFunObj[T any](value *T) FunObj[T] {
	return FunObj[T]{value}
}

// GetOrDefault Unwrapping function
func (s FunObj[T]) GetOrDefault(fun func() T) T {
	if s.value == nil {
		return fun()
	} else {
		return *s.value
	}
}

// GetOrNil Unwrapping function
func (s FunObj[T]) GetOrNil() *T {
	if s.value != nil {
		return s.value
	} else {
		return nil
	}
}
