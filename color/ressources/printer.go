package color

import (
	"fmt"
	"os"
	"strings"
)

var (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"
	Reset   = "\x1b[0m"
)

func Printer(inputLine string, slice [][]string, substring string, color string) {
	// Check for non-printable characters
	for _, char := range inputLine {
		if char < 32 || char > 126 {
			fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
			return
		}
	}
	for j := 0; j < 8; j++ {
		i := 0
		for i < len(inputLine) {
			if (len(os.Args) == 4 || len(os.Args) == 5) && strings.HasPrefix(inputLine[i:], substring) {
				for k := 0; k < len(substring); k++ {
					char := inputLine[i+k]
					index := int(char) - 32
					fmt.Print(getColor(color) + slice[index][j] + Reset)
				}
				i += len(substring) // skips the chars just processed
			} else {
				char := inputLine[i]
				index := int(char) - 32
				if len(os.Args) == 3 {
					fmt.Print(getColor(color) + slice[index][j])
				} else {
					fmt.Print(slice[index][j])
				}
				i++
			}
		}
		fmt.Println()
	}
}

func getColor(colors string) string {
	switch strings.ToLower(colors) {
	case "black":
		return Black
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	case "magenta":
		return Magenta
	case "cyan":
		return Cyan
	case "white":
		return White
	default:
		return ""
	}
}
