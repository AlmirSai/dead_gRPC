package main

import (
	"bytes"
	"io"
	"os"
	"log"
	"testing"
)

// TestMain captures the output of the main function and compares it to the expected output.
// It redirects standard output to a pipe, runs the main function, and then verifies if 
// the captured output matches the expected string. If it doesn't, the test reports an error.
func TestMain(t *testing.T) {

	oldStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	go func() {
		main()
		w.Close()
	}()

	var capturedOutput bytes.Buffer
	_, err = io.Copy(&capturedOutput, r)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = oldStdout

	log.Printf("Captured output: %q", capturedOutput.String())

	got := capturedOutput.String()
	expected := "Example of working with CI/CD\n"

	if got != expected {
		t.Errorf("main() = %v; want %v", got, expected)
	}
}