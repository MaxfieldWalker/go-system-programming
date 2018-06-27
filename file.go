package main

import (
	"os"
	"bytes"
	"fmt"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()

	os.Stdout.Write([]byte("os.Stdout example\n"))

	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
