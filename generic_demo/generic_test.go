package generic_demo

import "fmt"

func Get[T any](t T) T {
	switch v := any(&t).(type) {
	case *int:
		*v = 18
	}
	return t
}

func Get2[T any](t T) T {
	switch v := any(t).(type) {
	case int:
		v = 18
		fmt.Println(v)
	}
	return t
}
