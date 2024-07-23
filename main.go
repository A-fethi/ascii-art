package main

import (
	"fmt"
	"os"
	"strings"
)

var slice [][]string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_string>")
		return
	}

	content, err := os.ReadFile("/home/afethi/Desktop/ascii-art/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	str := string(content)
	lines := strings.Split(str, "\n\n")

	for i := range lines {
		if lines[i] != "" {
			if i == 0 {
				lines[0] = lines[0][1:]
			}
			slice = append(slice, strings.Split(lines[i], "\n"))
		}
	}
	input := os.Args[1]
	if input == "" {
		return
	} else if input == "\n" {
		print("\n")
	}
	inputLines := strings.Split(input, "\\n")
	for i, value := range inputLines {
		if value == "" {
			if i != 0 {
				fmt.Println()
			}
		} else {
			Printer(value)
		}
	}
}

func Printer(inputLine string) {
	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			if char >= 32 && char <= 126 { // printable ASCII range
				index := int(char) - 32
				if index >= 0 && index < len(slice) {
					fmt.Print(slice[index][j])
				} else {
					fmt.Printf("Character '%c' is out of range\n", char)
				}
			} else {
				fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			}
		}
		fmt.Println()
	}
}
