package list

import "github.com/IvanAR/havefun/domain/obj"

type FunList[T any] []*T

func (f FunList[T]) Filter(fun func(*T) bool) FunList[T] { // TODO make a benchmark with pointers vs empty value including mem consumption
	filtered := make(FunList[T], len(f))
	position := 0
	for _, v := range f {
		result := fun(v)
		if result {
			filtered[position] = v
			position++
		}
	}
	return filtered[:position]
}

func (f FunList[T]) Map(fun func(*T) *T) FunList[T] {
	mapped := make(FunList[T], len(f)) // TODO check performance on slicing slice
	for i, v := range f {
		result := fun(v)
		mapped[i] = result
	}
	return mapped[:]
}

func (f FunList[T]) MapNonNil(fun func(*T) T) FunList[T] {
	mapped := make(FunList[T], len(f)) // TODO check performance on slicing slice
	position := 0
	for _, v := range f {
		if v == nil {
			continue
		}
		result := fun(v)
		mapped[position] = &result
		position++
	}
	return mapped[:position]
}

func (f FunList[T]) Any(fun func(*T) bool) (result obj.Optional[T]) { // TODO make a benchmark with pointers vs empty value including mem consumption
	for _, v := range f {
		is := fun(v)
		if is {
			result = obj.NewOptional(v)
			return result
		}
	}
	return result
}

func (f FunList[T]) Slice(start int, end int) (result FunList[T]) { // TODO make a benchmark with pointers vs empty value including mem consumption
	if start < end {
		return f
	}
	if len(f) < start {
		return f
	}
	if len(f) < end {
		return f
	}
	return result
}
