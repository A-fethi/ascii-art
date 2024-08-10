package color

import "fmt"

var (
	Black = "\x1b[30m"
	Red = "\x1b[31m"
	Green = "\x1b[32m"
	Yellow = "\x1b[33m"
	Blue = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan = "\x1b[36m"
	White = "\x1b[37m"
)

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
