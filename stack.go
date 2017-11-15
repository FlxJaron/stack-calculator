package main

import "fmt"

// Arithmetic Operators
// https://golang.org/ref/spec#Arithmetic_operators
type ArithmeticOperator int

const (
	SUM ArithmeticOperator = iota
	DIFFERENCE
	PRODUCT
	QUOTIENT
)

// Stack properties
type Stack struct {
	values []float64
}

// Initialises a new Stack object with an empty map of values
func NewStack() *Stack {
	return &Stack{
		values: make([]float64, 0),
	}
}

// Push
func (stack *Stack) Push(number float64) {
	stack.values = append(stack.values, number)
}

// Pop
func (stack *Stack) Pop() (float64, error) {
	if len(stack.values) == 0 {
		return -1, fmt.Errorf("StackUnderflowException")
	}

	// Index
	i := len(stack.values) - 1

	// Read
	top := stack.values[i]

	// Remove
	stack.values = stack.values[:i]

	return top, nil
}

// Peek
func (stack *Stack) Peek() (float64, error) {
	if len(stack.values) == 0 {
		return -1, fmt.Errorf("StackUnderflowException")
	}

	return stack.values[len(stack.values)-1], nil
}

// Execute Arithmetic Operation
func (stack *Stack) ExecuteArithmeticOperation(arithmeticOperator ArithmeticOperator) (float64, error) {
	if len(stack.values) < 2 {
		return -1, fmt.Errorf("StackUnderflowException")
	}

	right, _ := stack.Pop()
	left, _ := stack.Pop()

	var result float64

	switch arithmeticOperator {
	case SUM:
		result = left + right
	case DIFFERENCE:
		result = left - right
	case PRODUCT:
		result = left * right
	case QUOTIENT:
		if right == 0 {
			return -1, fmt.Errorf("DivisionByZeroException")
		}
		result = left / right
	default:
		return -1, fmt.Errorf("Unsupported arithmetic operator: %s", arithmeticOperator)
	}

	// Only append to the stack when there is no error
	stack.Push(result)

	return result, nil
}
