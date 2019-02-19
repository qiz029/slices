package utils

import "reflect"

type (
	Utils interface {
		Contains_string(arr []string, ele string) (bool, error)
		Contains_int(arr []string, ele string) (bool, error)
	}

	Utilities struct {
	}
)

func (u *Utilities) Contains_string(arr []string, ele string) (bool, error) {
	for _, iter := range arr {
		if reflect.DeepEqual(iter, ele) {
			return true, nil
		}
	}
	return false, nil
}

func (u *Utilities) Contains_int(arr []int, ele int) (bool, error) {
	for _, iter := range arr {
		if reflect.DeepEqual(iter, ele) {
			return true, nil
		}
	}
	return false, nil
}
