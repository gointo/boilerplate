package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// Test Printhello returns values
func TestPrinthello(t *testing.T) {
	length, err := Printhello()
	if err != nil {
		t.Errorf("\nPrinthello returned an error: %v", err)
	} else if length != len("logger: YYYY/MM/DD HH/MM/SS hello world\n") {
		t.Errorf("\nPrinthello returned a wrong length: %v\nExpected: %v", length, len("logger: YYYY/MM/DD HH/MM/SS hello world\n"))
	}
}

// Check Output is ok
func TestPrinthelloOutput(t *testing.T) {
	backup := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	Printhello()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = backup
	if string(out) == "logger: YYYY/MM/DD HH/MM/SS hello world\n" {
		t.Errorf("\nPrinthello wrote a wrong string: %v\nExpected: %v", out, "logger: YYYY/MM/DD HH/MM/SS hello world\n")
	}
}

// Check Main function is ok
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
