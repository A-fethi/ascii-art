package main

import (
	"fmt"
	"os"
	"strings"
)
var sl  [][]string
func main() {
	// var result string
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_string>")
		return
	}

	content, err := os.ReadFile("/home/afethi/ascii-art/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	str := string(content)
	// input := os.Args[1]
	
	// result := []string{}
	lines := strings.Split(str, "\n\n")
	
	for i := range lines {
		if lines[i] != "" {
			if i == 0 {
				lines[0] = lines[0][1:]
			}
			sl = append(sl, strings.Split(lines[i], "\n"))
		}
	}
	// fmt.Println(sl[0])
	input := os.Args[1]
	sll := strings.Split(input, "\n")
	for _, v := range sll {
		Printer(v)
	}

	// fmt.Println(result)

	// fmt.Println(len(sl))
	// fmt.Println(sl[16])
}
func Printer(sll string ){
	j := 0
	pnl := ""
	for j < 8 {
		for _, char := range sll {
			
			if char >= 32 && char <= 126 { // printable ASCII range
				index := int(char) - 32
				if index >= 0 && index < len(sl) {
					// result = append(result, sl[index])
					fmt.Print(sl[index][j])

					// for i := 0; i < 8; i++ {
					// 	fmt.Print(sl[index][i])
					// }
				} else {
					fmt.Printf("Character '%c' is out of range\n", char)
				}
			} else if char == '/' || (char == 'n'&& pnl == "/") {
				pnl += string(char) 
				
				// fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			}
		}
		fmt.Println()

		j++
	}
}
