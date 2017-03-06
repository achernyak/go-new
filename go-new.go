package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s directory", os.Args[0])
		os.Exit(1)
	}
	dir := os.Args[1]

	var err error
	err = os.MkdirAll(dir, 0777)
	checkError(err)

	err = os.Chdir(dir)
	checkError(err)

	file := strings.Join([]string{dir, "go"}, ".")
	_, err = os.Create(file)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
