package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Size() int {
	return len(s.values)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Pop() (value T, success bool) {
	if s.IsEmpty() {
		return
	}

	idx := len(s.values) - 1
	value = s.values[idx]
	success = true
	s.values = s.values[:idx]

	return
}
