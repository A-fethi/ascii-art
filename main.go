package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var slice [][]string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_string>")
		return
	}
	content, err := os.ReadFile("standard.txt")
	if err != nil  || len(content) == 6623{
		url := "https://learn.zone01oujda.ma/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
		cmd := exec.Command("curl", "-o", "standard.txt", url)
		_, er := cmd.Output()
		if er != nil {
			fmt.Println("Error fetching url:", er)
			return
		}
		content, _ = os.ReadFile("standard.txt")
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
	} else {
		onlyNewLine := false

		inputLines := strings.Split(input, "\\n")
		for _, v := range inputLines {
			if v == "" {
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
				Printer(value)
			}
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
					return
				}
			} else {
				fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
				return
			}
		}
		fmt.Println()
	}
}
