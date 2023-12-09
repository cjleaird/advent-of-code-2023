package helper

import (
	"errors"
)

func Zip[T, S any](i1 []T, i2 []S) ([]Pair[T, S], error) {
	if len(i1) != len(i2) {
		return nil, errors.New("invalid inputs - lengths must be equal")
	}
	result := make([]Pair[T, S], len(i1))
	for i := 0; i < len(i1); i++ {
		result[i] = Pair[T, S]{i1[i], i2[i]}
	}
	return result, nil
}

type Pair[T, S any] struct {
	First  T
	Second S
}
