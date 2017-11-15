package main

import (
	"testing"
)

// Test Calculator Stack Pointers
func TestCalculatorStackPointers(t *testing.T) {
	calc := NewCalculator()

	// Get empty (new) stacks
	stack1 := calc.GetStack(1)
	stack2 := calc.GetStack(2)

	// Push and verify 123
	stack1.Push(123)
	stack1Peek, err := stack1.Peek()
	if err != nil || stack1Peek != 123 {
		t.Fail()
	}

	// Push and verify 456
	stack2.Push(456)
	stack2Peek, err := stack2.Peek()
	if err != nil || stack2Peek != 456 {
		t.Fail()
	}

	// Verify that the calculator returns the right stack pointers
	stack1PointerPeek, err := calc.GetStack(1).Peek()
	if err != nil || stack1PointerPeek != 123 {
		t.Fail()
	}

	stack2PointerPeek, err := calc.GetStack(2).Peek()
	if err != nil || stack2PointerPeek != 456 {
		t.Fail()
	}
}
