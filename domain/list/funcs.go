package list

//func Map[T interface{}, O interface{}](func(T) O) []T {
//
//}
//
//func List[T interface{}, O interface{}](func(T) O) []T {
//
//}

func Filter[T interface{}](values []*T, f func(*T) bool) []T {
	var filtered []T
	for _, v := range values {
		result := f(v)
		if result {
			filtered = append(filtered, *v)
		}
	}
	return filtered
}
