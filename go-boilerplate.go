package main

import (
	"bytes"
	"fmt"
	"log"
)

// Printhello argument on stdout
func Printhello() (int, error) {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.LstdFlags)
	logger.Print("hello world")
	return fmt.Print(&buf)
}

func main() {
	Printhello()
}
