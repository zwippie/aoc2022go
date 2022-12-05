package stack_test

import (
	"aoc2022/stack"
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack.Stack[int]{}
	if s.IsEmpty() == false {
		log.Fatal("isEmpty should be true")
	}
	s.Push(3)
	if s.IsEmpty() == true {
		log.Fatal("isEmpty should be false now")
	}
	s.Push(5)
	if s.Pop() != 5 {
		log.Fatal("wrong value popped")
	}
	if s.Pop() != 3 {
		log.Fatal("wrong value popped")
	}
}

// why do variables of this type have no
// access to the stack methods?
type ByteStack stack.Stack[byte]

func TestTypedStack(t *testing.T) {
	s := ByteStack{}
	if s.IsEmpty() == false {
		log.Fatal("isEmpty should be true")
	}
	s.Push(3)
	if s.IsEmpty() == true {
		log.Fatal("isEmpty should be false now")
	}
	s.Push(5)
	if s.Pop() != 5 {
		log.Fatal("wrong value popped")
	}
	if s.Pop() != 3 {
		log.Fatal("wrong value popped")
	}
}
