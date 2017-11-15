package main

import (
	"testing"
)

// Get Test Stack
func getTestStack() *Stack {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	return stack
}

// Test Peek
func TestPeek(t *testing.T) {
	peek, err := getTestStack().Peek()
	if err != nil || peek != 10 {
		t.Fail()
	}
}

// Test Push
func TestPush(t *testing.T) {
	stack := getTestStack()
	stack.Push(123)
	peek, err := stack.Peek()
	if err != nil || peek != 123 {
		t.Fail()
	}
}

// Test Pop
func TestPop(t *testing.T) {
	stack := getTestStack()
	stack.Pop()
	peek, err := stack.Peek()
	if err != nil || peek != 5 {
		t.Fail()
	}
}

// Test Arithmetic Operations
func TestArithmeticOperations(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result, err := getTestStack().ExecuteArithmeticOperation(SUM)
		if err != nil || result != 15 {
			t.Fail()
		}
	})
	t.Run("Subtract", func(t *testing.T) {
		result, err := getTestStack().ExecuteArithmeticOperation(DIFFERENCE)
		if err != nil || result != -5 {
			t.Fail()
		}
	})
	t.Run("Multiply", func(t *testing.T) {
		result, err := getTestStack().ExecuteArithmeticOperation(PRODUCT)
		if err != nil || result != 50 {
			t.Fail()
		}
	})
	t.Run("Divide", func(t *testing.T) {
		result, err := getTestStack().ExecuteArithmeticOperation(QUOTIENT)
		if err != nil || result != 0.5 {
			t.Fail()
		}
	})
}
