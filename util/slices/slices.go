package slices

import "reflect"

func ContainsAny(s []interface{}, t ...interface{}) bool {

	m := make(map[interface{}]bool)

	for _, v := range t {
		m[v] = true
	}

	for _, v := range s {
		if m[v] {
			return true
		}
	}
	return false
}

func Contains(s interface{}, t interface{}) bool {

	v := reflect.ValueOf(s)

	// 不是slice就抛出panic
	if v.Kind() != reflect.Slice {
		panic("input value is not slice")
	}

	// 目标类型
	targetValue := reflect.ValueOf(t)

	for i := 0; i < v.Len(); i++ {
		if targetValue == v.Index(i) {
			return true
		}
	}
	return false
}
