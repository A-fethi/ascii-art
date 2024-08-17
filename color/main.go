package main

import (
	"fmt"
	"os"
	"strings"

	color "color/ressources"
)

var colors string

// substring string

func formatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

func optionFlag() {
	option := os.Args[1:]
	for i := 0; i < len(option); i++ {
		if strings.HasPrefix(string(option[i]), "--color=") {
			colors = strings.TrimPrefix(string(option[i]), "--color=")
		}
	}
}

func main() {
	var banner string
	var input string
	var substring string

	if len(os.Args) == 2 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			formatError()
			return
		} else {
			input = os.Args[1]
			banner = "standard"
		}
	} else if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			input = os.Args[2]
			banner = "standard"
			optionFlag()
		} else {
			input = os.Args[1]
			banner = os.Args[2]
		}
	} else if len(os.Args) == 4 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			substring = os.Args[2]
			input = os.Args[3]
			banner = "standard"
			optionFlag()
		} else {
			formatError()
			return
		}
	} else if len(os.Args) == 5 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			substring = os.Args[2]
			input = os.Args[3]
			banner = os.Args[4]
			optionFlag()
		} else {
			formatError()
			return
		}
	} else {
		formatError()
		return
	}
	if !strings.HasPrefix(os.Args[1], "--color=") {
		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			fmt.Println("Error: Not a valid banner")
			return
		}
	}
	str := color.Reader(banner)
	color.Splitter(banner, str)
	if input == "" {
		return
	} else {
		inputLines := strings.Split(input, "\\n")
		// checking if the input has only newlines
		inputLines = color.OnlyNewLine(inputLines)
		for _, value := range inputLines {
			if value == "" {
				fmt.Println()
			} else {
				color.Printer(value, color.Slice, substring, colors)
			}
		}
	}
}
