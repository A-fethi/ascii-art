package main

import (
	"fmt"
	"os"
	"strings"

	fs "fs/ressources"
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
	if input == "" {
		return
	} else {
		// gets rid of the first empty string if the input has ONLY newlines
		onlyNewLine := false
		inputLines := strings.Split(input, "\\n")
		for _, value := range inputLines {
			if value == "" {
				onlyNewLine = true
			} else {
				onlyNewLine = false
				break
			}
		}
		if onlyNewLine {
			inputLines = inputLines[1:]
		}
		for _, value := range inputLines {
			if value == "" {
				fmt.Println()
			} else {
				fs.Printer(value, slice)
			}
		}
	}
}
