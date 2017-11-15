package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Calculator properties
type Calculator struct {
	stacks map[int]*Stack
	Router *mux.Router
}

// Initialises a new Calculator object with an empty map of stacks
func NewCalculator() *Calculator {
	calc := &Calculator{
		stacks: make(map[int]*Stack, 0),
		Router: mux.NewRouter().PathPrefix("/calc/{id:[0-9]+}").Subrouter(),
	}

	// Initialise router with paths and handlers
	calc.Router.HandleFunc("/peek", calc.HandlePeek)
	calc.Router.HandleFunc("/push/{number}", calc.HandlePush)
	calc.Router.HandleFunc("/pop", calc.HandlePop)
	calc.Router.HandleFunc("/add", calc.HandleAdd)
	calc.Router.HandleFunc("/subtract", calc.HandleSubtract)
	calc.Router.HandleFunc("/multiply", calc.HandleMultiply)
	calc.Router.HandleFunc("/divide", calc.HandleDivide)
	calc.Router.NotFoundHandler = http.HandlerFunc(calc.HandleNotFound)

	return calc
}

// Get Stack by id
// returns a new object when the id is not found in the map
func (calc *Calculator) GetStack(id int) *Stack {
	if calc.stacks[id] == nil {
		calc.stacks[id] = NewStack()
	}
	return calc.stacks[id]
}

// Get Stack From Request
func (calc *Calculator) getStackFromRequest(r *http.Request) (*Stack, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	return calc.GetStack(id), err
}

// Returns stack[top]
func (calc *Calculator) HandlePeek(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		result, err := stack.Peek()
		if err != nil {
			calc.sendResponse(w, err.Error())
			return
		}
		calc.sendResponse(w, fmt.Sprintf("%g", result))
	}
}

// Pushes a number onto the stack
func (calc *Calculator) HandlePush(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		vars := mux.Vars(r)
		number, err := strconv.ParseFloat(vars["number"], 64)
		if err != nil {
			calc.sendResponse(w, "InvalidNumberException")
			return
		}
		stack.Push(number)
		calc.sendResponse(w, fmt.Sprintf("%g", number))
	}
}

// Returns the top from the stack and removes it
func (calc *Calculator) HandlePop(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		stack.Pop()
	}
}

// Removes the top and top-1 from the stack and replaces it with stack[top-1]+stack[top]
func (calc *Calculator) HandleAdd(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		result, _ := stack.ExecuteArithmeticOperation(SUM)
		if err != nil {
			calc.sendResponse(w, err.Error())
			return
		}
		calc.sendResponse(w, fmt.Sprintf("%g", result))
	}
}

// Removes the top and top-1 from the stack and replaces it with stack[top-1]-stack[top]
func (calc *Calculator) HandleSubtract(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		result, err := stack.ExecuteArithmeticOperation(DIFFERENCE)
		if err != nil {
			calc.sendResponse(w, err.Error())
			return
		}
		calc.sendResponse(w, fmt.Sprintf("%g", result))
	}
}

// Removes the top and top-1 from the stack and replaces it with stack[top-1]*stack[top]
func (calc *Calculator) HandleMultiply(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		result, err := stack.ExecuteArithmeticOperation(PRODUCT)
		if err != nil {
			calc.sendResponse(w, err.Error())
			return
		}
		calc.sendResponse(w, fmt.Sprintf("%g", result))
	}
}

// Removes the top and top-1 from the stack and replaces it with stack[top-1]/stack[top]
func (calc *Calculator) HandleDivide(w http.ResponseWriter, r *http.Request) {
	if stack, err := calc.getStackFromRequest(r); err == nil {
		result, err := stack.ExecuteArithmeticOperation(QUOTIENT)
		if err != nil {
			calc.sendResponse(w, err.Error())
			return
		}
		calc.sendResponse(w, fmt.Sprintf("%g", result))
	}
}

// Handle Not Found
// e.g. when calling the calc with an invalid stack id: /calc/abc/peek
func (calc *Calculator) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	calc.sendResponse(w, "Please read the documentation for finding the appropriate method")
}

// Send response as plain text
func (calc *Calculator) sendResponse(w http.ResponseWriter, response string) {
	w.Write([]byte(response))
}
