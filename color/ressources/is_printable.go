package color

import "fmt"

func IsPrintable(input string) bool {
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			return false
		}
	}
	return true
}
