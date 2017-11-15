# stack-calculator
Code assignment for a Stack Calculator in Go

## Build
1) export GOPATH=\`pwd\`
2) go get github.com/gorilla/mux
3) `go build .`

The mux dependency is necessary for dynamic parameters in the router paths like `{id}` and `{number}`

## Test (verbose)
Unit test and integration test can be executed with `go test -v`

## Run
To run the application, you can execute the binary after following the build steps.

`./stack-calculator`

## Endpoints
All HTTP endpoints listen on GET requests by default on port `8080`.
The endpoints provide operations that work on a set of stack calculator. The parameter `:id` must be an integer (`[0-9]+`) at all times.

#### `/calc/{id}/peek` 
- Returns stack[top]
- Returns `StackUnderflow` when the stack contains no values

#### `/calc/{id}/push/{number}`
- Pushes a number onto the stack
- Returns the new stack[top]
- Returns `InvalidNumberException` when `{number}` is not a valid float
 
#### `/calc/:id/pop`
- Returns the top from the stack and removes it
- Returns `StackUnderflow` when the stack is empty
 
#### `/calc/:id/add`
 - Removes the top and top-1 from the stack and replaces it with stack[top-1]+stack[top]
 - Returns the new stack[top]
 - Returns `StackUnderflow` when the stack contains less than two values
 
#### `/calc/:id/subtract`
 - Removes the top and top-1 from the stack and replaces it with stack[top-1]-stack[top]
 - Returns the new stack[top]
 - Returns `StackUnderflow` when the stack contains less than two values
 
#### `/calc/:id/multiply`
 - Removes the top and top-1 from the stack and replaces it with stack[top-1]*stack[top]
 - Returns the new stack[top]
 - Returns `StackUnderflow` when the stack contains less than two values
 
#### `/calc/:id/divide`
 - Removes the top and top-1 from the stack and replaces it with stack[top-1]/stack[top]
 - Returns the new stack[top]
 - Returns `StackUnderflow` when the stack contains less than two values
 - Returns `DevisionByZero` when stack[top] is `0` 

### Improvements
 - Make the assignment less detailed to make room for creativity
 - Continuous Integration tool
 - Run application in docker container
 - Shell scripts for building, running and testing
 - Stateless (RESTful) API
 - Error structs
 - Configuration file (e.g. server port)
 - Logging (e.g. verbose/external for debug purposes)