package main

import "testing"

func TestOutput(t *testing.T) {

	if expected, actual := "Hello World", messageOutput("Hello World"); actual != expected {
		t.Errorf("The main function provided an actual of %q but %q was expected.", actual, expected)
	}
}
