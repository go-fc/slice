package slice

import "errors"

var ErrNotFound = errors.New("not found")

func Map[P any, R any](data []P, fn func(P) R) (r []R) {
	for _, item := range data {
		r = append(r, fn(item))
	}
	return
}

func Filter[P any](data []P, fn func(P) bool) (r []P) {
	for _, item := range data {
		if fn(item) {
			r = append(r, item)
		}
	}
	return
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
	return r, ErrNotFound
}

func Copy[T any](data []T) (r []T) {
	return append(r, data...)
}

func ToAny[T any](list []T) (r []any) {
	for _, item := range list {
		r = append(r, item)
	}
	return
}

func Safe[T any](data []T, index int) (val T) {
	if len(data) > index {
		return data[index]
	}
	return
}
