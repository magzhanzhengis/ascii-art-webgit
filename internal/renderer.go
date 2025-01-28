package internal

import (
	"fmt"
	"strings"
)

// RenderASCIIArt converts text to ASCII using the given templates.
func RenderASCIIArt(text string, asciiTemplates [][]string) string {
	var result string
	for i := 0; i < len(asciiTemplates[0]); i++ { // Loop over rows of ASCII characters
		for _, char := range text {
			if int(char)-32 >= 0 && int(char)-32 < len(asciiTemplates) {
				result += asciiTemplates[int(char)-32][i] + " "
			}
		}
		result += "\n"
	}
	return result
}

// RenderASCIIArt processes the input string and prints its ASCII art representation.
// func RenderASCIIArt(text string, asciiTemplates [][]string) {
// 	// Split input text by newline markers
// 	substrings := SplitNewline(text)

// 	// Variable to track the count of consecutive newlines
// 	wordnumber := 0
// 	wordcount := 0
// 	for i := 0; i < len(substrings); i++ {
// 		if substrings[i] != "\\n" {
// 			wordnumber++
// 		}
// 	}
// 	for i := 0; i < len(substrings); i++ {
// 		if wordcount == 0 && substrings[i] == "\\n" {
// 			fmt.Println()
// 		} else if wordcount != 0 && substrings[i] == "\\n" && wordcount < wordnumber {
// 			if substrings[i+1] != "\\n" {
// 				continue
// 			}
// 			fmt.Println()
// 		} else if wordcount == wordnumber && substrings[i] == "\\n" {
// 			fmt.Println()
// 		} else {
// 			PrintASCIICharacters(substrings[i], asciiTemplates)
// 			wordcount++
// 		}
// 	}
// }

// PrintASCIICharacters prints the ASCII art representation of a given string.
func PrintASCIICharacters(text string, asciiTemplates [][]string) {
	// Convert the text into indices corresponding to ASCII characters
	charIndices := ConvertTextToASCIIIndices(text)
	// Render each line of the ASCII art (8 lines per character)
	for i := 0; i < 8; i++ {
		var line strings.Builder // Efficient string concatenation
		for j, index := range charIndices {
			if index >= 0 && index < len(asciiTemplates) {
				line.WriteString(asciiTemplates[index][i]) // Append ASCII art for the character
				// } else {
				// 	line.WriteString("[, ], !") // Handle unsupported characters
				// }
				// Add a space between characters except the last one
				if j < len(charIndices)-1 {
					line.WriteString(" ")
				}
			} // Print the constructed line
		}
		fmt.Println(line.String())
	}
}
