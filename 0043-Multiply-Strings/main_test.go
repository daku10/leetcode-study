package main

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {

	type TestCase struct {
		argNum1 string
		argNum2 string
		expect string
	}

	testCases := []TestCase{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"123456789123456789123456789123456789123456789123456789123456789", "123456789123456789123456789123456789123456789123456789123456789123456789", "15241578780673678546105778311537878076969977842402077607834177358024698342783119577351019811918920046486820281054720515622620750190521"},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			actual := multiply(testCase.argNum1, testCase.argNum2)
			if actual != testCase.expect {
				t.Errorf("num1 :%s num2: %s expect: %s actual :%s", testCase.argNum1, testCase.argNum2, testCase.expect, actual)
			}
		})
	}    
}
