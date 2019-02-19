package slices

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/qiz029/slices/utils"
)

type Slices struct {
	u *utils.Utilities
}

func (s Slices) Contains(arr interface{}, ele interface{}) (bool, error) {
	rt := reflect.TypeOf(arr)
	// err := errors.New(fmt.Sprintf("is %v and is %v", rt.Elem(), reflect.TypeOf(ele)))
	// if err != nil {
	// 	return false, err
	// }
	var err error
	if rt.Kind() != reflect.Slice && rt.Kind() != reflect.Array {
		err = errors.New("The first argument is not a slice or an array")
		return false, err
	}
	typeOfElement := reflect.TypeOf(ele)
	if typeOfElement != rt.Elem() {
		err = errors.New("The type of the element in slice is not the same as the type of the element")
		return false, err
	}

	switch ele.(type) {
	case int:
		arr1 := arr.([]int)
		return s.u.Contains_int(arr1, ele.(int))
	case string:
		arr1 := arr.([]string)
		return s.u.Contains_string(arr1, ele.(string))
	default:
		break
	}
	return false, errors.New(fmt.Sprintf("cannot handle this type %s", typeOfElement))

}
