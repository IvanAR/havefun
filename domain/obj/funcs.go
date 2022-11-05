package obj

type Optional[T any] struct {
	value *T
}

func NewOptional[T any](value *T) Optional[T] {
	return Optional[T]{value}
}

// GetOrDefault Unwrapping function
func (s Optional[T]) GetOrDefault(fun func() T) T {
	if s.value == nil {
		return fun()
	} else {
		return *s.value
	}
}

// GetOrNil Unwrapping function
func (s Optional[T]) GetOrNil() *T {
	if s.value != nil {
		return s.value
	} else {
		return nil
	}
}
