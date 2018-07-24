package go_utility

import "testing"

func TestStack_Clear(t *testing.T) {
	s := Stack{}
	AssertTrue(t, s.Size() == 0)
	s.Push("One")
	s.Push("Two")
	s.Push("Three")
	s.Push("Four")
	AssertTrue(t, s.Size() == 4)
	s.Clear()
	AssertTrue(t, s.Size() == 0)
}

func TestStack_Size(t *testing.T) {
	s := Stack{}
	AssertTrue(t, s.Size() == 0)
	s.Push("One")
	AssertTrue(t, s.Size() == 1)
	s.Push("Two")
	AssertTrue(t, s.Size() == 2)
	s.Push("Three")
	AssertTrue(t, s.Size() == 3)
	s.Push("Four")
	AssertTrue(t, s.Size() == 4)
}

func TestStack_Peek(t *testing.T) {
	s := Stack{}
	AssertNil(t, s.Peek())
	s.Push("One")
	AssertEquals(t, "One", s.Peek())
	s.Push("Two")
	AssertEquals(t, "Two", s.Peek())
	s.Push("Three")
	AssertEquals(t, "Three", s.Peek())
	s.Push("Four")
	AssertEquals(t, "Four", s.Peek())
}

func TestStack_Pop(t *testing.T) {
	s := Stack{}
	AssertNil(t, s.Pop())
	s.Push("One")
	s.Push("Two")
	s.Push("Three")
	s.Push("Four")
	AssertEquals(t, "Four", s.Pop())
	AssertEquals(t, "Three", s.Pop())
	AssertEquals(t, "Two", s.Pop())
	AssertEquals(t, "One", s.Pop())
	AssertNil(t, s.Pop())
}
