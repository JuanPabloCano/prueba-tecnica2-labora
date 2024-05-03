package common

import "math/rand"

func GetRandomObject[T any](slice []T) T {
	temp := make([]T, len(slice))
	copy(temp, slice)
	idx := rand.Intn(len(temp))
	return slice[idx]
}
