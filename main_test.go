package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

// Test Case
type testCase struct {
	endpoint      string
	expectedValue string
}

// Test Integration
func TestIntegration(t *testing.T) {
	// Initialise New Calculator
	calc := NewCalculator()

	// Http Test Server
	server := httptest.NewServer(calc.Router)

	// Test cases map with indexes for making sort possible, because of not sorted hash tables
	tests := map[int]*testCase{
		0: &testCase{
			endpoint:      "1/push/1",
			expectedValue: "1",
		},
		1: &testCase{
			endpoint:      "1/push/4",
			expectedValue: "4",
		},
		2: &testCase{
			endpoint:      "1/peek",
			expectedValue: "4",
		},
		3: &testCase{
			endpoint:      "1/add",
			expectedValue: "5",
		},
		4: &testCase{
			endpoint:      "1/push/10",
			expectedValue: "10",
		},
		5: &testCase{
			endpoint:      "1/multiply",
			expectedValue: "50",
		},
		6: &testCase{
			endpoint:      "1/push/2",
			expectedValue: "2",
		},
		7: &testCase{
			endpoint:      "1/divide",
			expectedValue: "25",
		},
		8: &testCase{
			endpoint:      "2/peek",
			expectedValue: "StackUnderflowException",
		},
		9: &testCase{
			endpoint:      "2/push/20",
			expectedValue: "20",
		},
		10: &testCase{
			endpoint:      "2/divide",
			expectedValue: "StackUnderflowException",
		},
		11: &testCase{
			endpoint:      "2/push/0",
			expectedValue: "0",
		},
		12: &testCase{
			endpoint:      "2/push/0",
			expectedValue: "0",
		},
		13: &testCase{
			endpoint:      "2/divide",
			expectedValue: "DivisionByZeroException",
		},
		14: &testCase{
			endpoint:      "invalid-endpoint",
			expectedValue: "Please read the documentation for finding the appropriate method",
		},
		15: &testCase{
			endpoint:      "1/push/invalid-number",
			expectedValue: "InvalidNumberException",
		},
	}

	// Sort by key
	var keys []int
	for k := range tests {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Run test cases in chronological order
	// when the result is not identical to the expected value, the test fails
	for _, k := range keys {
		testCase := tests[keys[k]]
		t.Run(testCase.endpoint, func(t *testing.T) {
			url := fmt.Sprintf("%s/calc/%s", server.URL, testCase.endpoint)
			if httpGet(url) != testCase.expectedValue {
				fmt.Println(url, testCase.expectedValue)
				t.Fail()
			}
		})
	}

	defer server.Close()
}

// Execute HTTP Call on the http test server
func httpGet(url string) string {
	r, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	responseBytes, _ := ioutil.ReadAll(r.Body)
	x := string(responseBytes[:])
	return x
}
