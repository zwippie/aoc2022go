package stack

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// func Push[S ~[]T, T any](s []T, v T) []T {
// 	return append(s, v)
// }

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// func Pop[T any](s []T) (T, []T) {
// 	return s[len(s)-1], s[:len(s)-1]
// }

func (s *Stack[T]) Pop() T {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
