package stack_test

import (
	"aoc2022/stack"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack.Stack[int]{}
	if s.IsEmpty() == false {
		t.Error("stack should be empty")
	}
	s.Push(3)
	if s.IsEmpty() == true {
		t.Error("stack should not be empty")
	}
	s.Push(5)
	got := s.Pop()
	if got != 5 {
		t.Errorf("s.Pop() = %d; want 5", got)
	}
	if s.Pop() != 3 {
		t.Errorf("s.Pop() = %d; want 3", got)
	}
}

// why do variables of this type have no
// access to the stack methods?
type ByteStack stack.Stack[byte]

// func TestTypedStack(t *testing.T) {
// 	s := ByteStack{}
// 	if s.IsEmpty() == false {
// 		t.Error("stack should be empty")
// 	}
// 	s.Push(3)
// 	if s.IsEmpty() == true {
// 		t.Error("stack should not be empty")
// 	}
// 	s.Push(5)
// 	got := s.Pop()
// 	if got != 5 {
// 		t.Errorf("s.Pop() = %d; want 5", got)
// 	}
// 	if s.Pop() != 3 {
// 		t.Errorf("s.Pop() = %d; want 3", got)
// 	}
// }
