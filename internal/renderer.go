package internal

import (
	"strings"
)

// RenderASCIIArt processes the input string and returns its ASCII art representation as a string.
func RenderASCIIArt(text string, asciiTemplates [][]string) string {
	// Split input text by newline markers
	substrings := SplitNewline(text)
	// Variable to track the count of consecutive newlines
	var result strings.Builder
	// Loop through the substrings and process each one
	for i := 0; i < len(substrings); i++ {
		if substrings[i] == "\\n" {
			result.WriteString("\n") // Handle escaped newlines
		} else {
			// Process each word individually
			result.WriteString(PrintASCIICharacters(substrings[i], asciiTemplates))
		}
	}
	return result.String()
}

// PrintASCIICharacters generates the ASCII art for each character in the input text.
func PrintASCIICharacters(text string, asciiTemplates [][]string) string {
	// Convert the text into indices corresponding to ASCII characters
	charIndices := ConvertTextToASCIIIndices(text)
	// Prepare a string builder to accumulate the result
	var result strings.Builder
	// Render each line of the ASCII art (assuming 8 lines per character)
	for i := 0; i < 8; i++ {
		for j, index := range charIndices {
			// Ensure valid index and print ASCII art for each character
			if index >= 0 && index < len(asciiTemplates) {
				result.WriteString(asciiTemplates[index][i])
			}
			// Add a space between characters except the last one
			if j < len(charIndices)-1 {
				result.WriteString(" ")
			}
		}
		// Add a newline after each row of characters
		result.WriteString("\n")
	}
	return result.String()
}
