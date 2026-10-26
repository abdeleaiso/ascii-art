package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// project accept only 2 args
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . {string}")
		return
	}

	input := os.Args[1]
	if input == "" {
		return
	}

	// read banner file he return []byte
	fileData, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// in windows they have \r\n intead of \n
	fileText := strings.ReplaceAll(string(fileData), "\r\n", "\n")

	// split string with \n into slice of strings
	lines := strings.Split(fileText, "\n")

	result := TransformToAscii(input, lines)

	fmt.Print(result)
}
