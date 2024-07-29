package main

import (
	"fmt"
	"os"
	"strings"
)

var slice [][]string

func main() {
	var banner string
	var input string

	if len(os.Args) == 2 {
		input = os.Args[1]
		banner = "standard"
	} else if len(os.Args) == 3 {
		input = os.Args[1]
		banner = os.Args[2]
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . something standard")
		return
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("Error: Not a valid banner")
		return
	}
	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	str := string(content)
	var lines []string
	if banner != "thinkertoy" {
		lines = strings.Split(str, "\n\n")
	} else {
		fixedContent := strings.ReplaceAll(str, "\r\n", "\n")
		lines = strings.Split(fixedContent, "\n\n")
	}

	for i := range lines {
		if lines[i] != "" {
			if i == 0 {
				lines[0] = lines[0][1:]
			}
			slice = append(slice, strings.Split(lines[i], "\n"))
		}
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
				os.Exit(1)
			}
		}
		fmt.Println()
	}
}


/*check if valid charsare given befor printing 
and pro*/