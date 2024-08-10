package color

// gets rid of the first empty string if the input has ONLY newlines
func OnlyNewLine(inputLines []string) []string {
	onlyNewLine := false
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
	return inputLines
}
