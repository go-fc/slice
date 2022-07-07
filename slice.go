package slice

import "errors"

func Map[P any, R any](data []P, fn func(P) R) []R {
	result := make([]R, 0, len(data))
	for _, item := range data {
		result = append(result, fn(item))
	}
	return result
}

func Filter[P any](data []P, fn func(P) bool) []P {
	result := make([]P, 0)
	for _, item := range data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Contains[T comparable](data []T, e T) bool {
	for _, item := range data {
		if item == e {
			return true
		}
	}
	return false
}

func Find[T any](data []T, fn func(T) bool) (r T, err error) {
	for _, item := range data {
		if fn(item) {
			return item, nil
		}
	}
	return r, errors.New("not found")
}

func Copy[T any](data []T) []T {
	return append(make([]T, 0, len(data)), data...)
}
