package internal

import (
	"fmt"
	"os"
	"strings"
)

// ReadBannerFile reads the content of the banner file and returns it as a slice of strings.
func ReadBannerFile(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read banner file %s: %v", filePath, err))
	}
	return strings.Split(string(content), "\n")
}

// ParseBanner converts banner file lines into a 2D slice representing ASCII templates.
func ParseBanner(fileLines []string) [][]string {
	var asciiTemplates [][]string
	var tempASCII []string
	counter := 0

	for _, line := range fileLines {
		counter++
		if counter != 1 {
			tempASCII = append(tempASCII, line)
		}
		if counter == 9 {
			asciiTemplates = append(asciiTemplates, tempASCII)
			counter = 0
			tempASCII = nil
		}
	}

	return asciiTemplates
}
