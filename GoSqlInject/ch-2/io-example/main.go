package main

import (
	"fmt"
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

	input := make([]byte, 4096)
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = wtiter.Write(input)
	if err != nil {
		log.Fatalf("Unable to wirte data:%v", err)
	}
	fmt.Printf("Wrote %d bytes to student\n", s)
}
