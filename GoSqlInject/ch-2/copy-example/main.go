package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct {
}

func (fooreader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

type FooWrite struct {
}

func (fooWrite *FooWrite) Write(b []byte) (int, error) {
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		wtiter FooWrite
	)

	if _, err := io.Copy(&wtiter, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
