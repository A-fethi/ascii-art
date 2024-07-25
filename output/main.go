package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	slice      [][]string
	outputFile string
)

func formatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func optionFlag() {
	option := os.Args[1:]
	for i := 0; i < len(option); i++ {
		if strings.HasPrefix(string(option[i]), "--output=") {
			outputFile = strings.TrimPrefix(string(option[i]), "--output=")
		}
	}
}

func Printer(inputLine string) string {
	result := ""
	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			if char >= 32 && char <= 126 { // printable ASCII range
				index := int(char) - 32
				if index >= 0 && index < len(slice) {
					if len(os.Args) < 3 {
						fmt.Print(slice[index][j])
					} else {
						result += slice[index][j]
					}
				} else {
					fmt.Printf("Character '%c' is out of range\n", char)
				}
			} else {
				fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			}
		}
		if len(os.Args) < 3 {
			fmt.Println()
		} else {
			result += "\n"
		}
	}
	return result
}

func main() {
	var banner string
	var input string

	if len(os.Args) == 2 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			formatError()
			return
		} else {
			input = os.Args[1]
			banner = "standard"
		}
	} else if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			input = os.Args[2]
			banner = "standard"
			optionFlag()
		} else {
			input = os.Args[1]
			banner = os.Args[2]
		}
	} else if len(os.Args) == 4 {
		input = os.Args[2]
		banner = os.Args[3]
		optionFlag()
	} else {
		formatError()
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
	var fullOutput string
	for i, value := range inputLines {
		if value == "" {
			if i != 0 {
				fullOutput += "\n"
			}
		} else {
			output := Printer(value)
			fullOutput += output
		}
	}
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(fullOutput), 0o644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
