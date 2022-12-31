package helper

func Map[T, R any](source []T, transform func(T) R) []R {
	res := make([]R, 0, len(source))
	for _, item := range source {
		res = append(res, transform(item))
	}
	return res
}

func Filter[T any](source []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(source))
	for _, item := range source {
		if predicate(item) {
			res = append(res, item)
		}
	}
	return res
}
