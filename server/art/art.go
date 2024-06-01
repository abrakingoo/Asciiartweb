package art

import (
	"fmt"
	"os"
	"strings"
)

// GetStartingIndex calculates the starting index of the ASCII character in the file
func GetStartingIndex(ascii int) int {
	return (ascii-32)*9 + 1
}

// File returns the path to the ASCII art file based on the banner name
func File(name string) string {
	files := map[string]string{
		"standard":   "server/art/standard.txt",
		"thinkertoy": "server/art/thinkertoy.txt",
		"shadow":     "server/art/shadow.txt",
	}

	if file, exists := files[name]; exists {
		return file
	}
	return "No Match Found For The File " + name
}

// handleBackspace processes backspace characters in the input string
func handleBackspace(s string) string {
	index := strings.Index(s, "\\b")

	if index == -1 {
		return s
	}
	if index == 0 {
		return handleBackspace(s[1:])
	}
	return handleBackspace(s[:index-1] + s[index+2:])
}

// Art generates the ASCII art based on the input string and fileName
func Art(input, fileName string) string {
	if input == "\\n" {
		fmt.Println()
		os.Exit(0)
	}
	input = handleBackspace(input)
	input = strings.ReplaceAll(string(input), "\\t", "   ")
	file, err := os.ReadFile(File(fileName))

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	fileData := strings.Split(string(file), "\n")

	var outputBuilder strings.Builder
	inputParts := strings.Split(input, "\\n")

	for _, part := range inputParts {
		if part == "" {
			outputBuilder.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range part {
				StartingIndex := GetStartingIndex(int(char))
				outputBuilder.WriteString(fileData[StartingIndex+i])
			}
			outputBuilder.WriteString("\n")
		}
	}
	return outputBuilder.String()
}
