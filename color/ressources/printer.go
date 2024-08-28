package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	Black   = "\x1b[38;2;0;0;0m"
	Red     = "\x1b[38;2;255;0;0m"
	Green   = "\x1b[38;2;0;255;0m"
	Yellow  = "\x1b[38;2;255;255;0m"
	Blue    = "\x1b[38;2;0;0;255m"
	Magenta = "\x1b[38;255;0;0;255m"
	Cyan    = "\x1b[38;2;43;255;255m"
	White   = "\x1b[38;2;255;255;255m"
	Reset   = "\x1b[0m"
)

func Printer(inputLine string, slice [][]string, substring string, color string) {
	// Check for non-printable characters
	for j := 0; j < 8; j++ {
		i := 0
		for i < len(inputLine) {
			if (len(os.Args) == 4 || len(os.Args) == 5) && strings.HasPrefix(inputLine[i:], substring) {
				for k := 0; k < len(substring); k++ {
					char := inputLine[i+k]
					index := int(char) - 32
					fmt.Print(GetColor(color) + slice[index][j] + Reset)
				}
				i += len(substring) // skips the chars just processed
			} else {
				char := inputLine[i]
				index := int(char) - 32
				if len(os.Args) == 3 {
					fmt.Print(GetColor(color) + slice[index][j])
				} else {
					fmt.Print(slice[index][j])
				}
				i++
			}
		}
		fmt.Println()
	}
}

func GetColor(colors string) string {
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
		if strings.HasPrefix(colors, "rgb(") && strings.HasSuffix(colors, ")") {
			rgbRegex := regexp.MustCompile(`rgb\((\d+), (\d+), (\d+)\)`)
			match := rgbRegex.FindStringSubmatch(colors)
			if match != nil {
				r, _ := strconv.Atoi(match[1])
				g, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
			}
		} else if strings.HasPrefix(colors, "#") && len(colors) == 7 {

			r, err := strconv.ParseInt(colors[0:2], 16, 8)
			g, err1 := strconv.ParseInt(colors[2:4], 16, 8)
			b, err2 := strconv.ParseInt(colors[4:], 16, 8)

			if err != nil || err1 != nil || err2 != nil {
				fmt.Println("ERROR: the color isn't available.\nRetry with a valid rgb or hex code.")
				fmt.Println()
				os.Exit(69)
			}

			return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)

		}
		return ""
	}
} //#FF0000
