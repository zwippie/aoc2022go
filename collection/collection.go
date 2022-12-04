package collection

import "fmt"

// type Collection []int

type Element interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[V Element](values []V) V {
	var sum V
	for _, value := range values {
		sum += value
	}
	return sum
}

func Last[V Element](values []V) V {
	return values[len(values)-1]
}

func LastN[V Element](values []V, n int) []V {
	return values[len(values)-n:]
}

func Map[V Element](values []V, fn func(V) any) []any {
	result := []any{}
	for _, value := range values {
		result = append(result, fn(value))
	}
	return result
}

func SliceMap[K, V comparable](input []K, f func(K) V) []V {
	result := []V{}
	for _, v := range input {
		result = append(result, f(v))

	}
	return result
}

// func (c Collection) Last() int {
// 	return c[len(c)-1]
// }

// func (c Collection) LastN(n int) Collection {
// 	return c[len(c)-n:]
// }

func Test() {
	fmt.Println("testing 123")
}
