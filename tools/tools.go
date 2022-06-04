package tools

import "fmt"

func Remove[T any](slice []T, index int) ([]T, error) {
	if len(slice) <= index {
		return slice, fmt.Errorf("index out of bound")
	}

	slice = append(slice[:index], slice[index+1:]...)
	return slice, nil
}
