package fs

import "fmt"

func Printer(inputLine string, slice [][]string) {
	// checks if the char is not printable
	for _, char := range inputLine {
		if char < 32 || char > 126 {
			fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			return
		}
	}

	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			fmt.Print(slice[index][j])
		}
		fmt.Println()
	}
}
