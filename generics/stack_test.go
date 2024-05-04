package generics

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("General operations on stack", func(t *testing.T) {
		var s Stack[int]
		s.Push(3)
		s.Push(4)

		if s.Size() != 2 {
			t.Errorf("Expected 2 items, but got %d items", s.Size())
		}

		r, success := s.Pop()

		if !success {
			t.Errorf("Expected to success but failed")
		}

		if r != 4 {
			t.Errorf("Expected 4, but got %d", r)
		}
	})

	t.Run("Pop on empty stack", func(t *testing.T) {
		var s Stack[int]
		ret, success := s.Pop()

		AssertEqual(t, ret, 0)
		AssertEqual(t, success, false)
	})
}
