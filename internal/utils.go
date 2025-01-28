package internal

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode"
)

// SplitTextByNewline splits the input text into segments based on escaped newlines.
func SplitNewline(text string) []string {
	var lines []string

	// Compile the regular expression to match both \n and \\n
	re := regexp.MustCompile(`[\n]|\\n`)

	// Split the text based on the regex
	parts := re.Split(text, -1)

	// Find all the newline delimiters (\n or \\n) to keep them in the result
	delimiters := re.FindAllString(text, -1)

	// Combine the split parts and delimiters into the final result
	for i, part := range parts {
		if part != "" {
			lines = append(lines, part) // Add the text segment
		}
		if i < len(delimiters) {
			if delimiters[i] == "\n" {
				lines = append(lines, "\\n")
			} else {
				lines = append(lines, delimiters[i]) // Add the corresponding delimiter
			}

		}
	}

	return lines
}

// ConvertTextToASCIIIndices converts a string into ASCII indices for the banner.
func ConvertTextToASCIIIndices(text string) []int {
	var indices []int
	for _, char := range text {
		indices = append(indices, int(char)-32) // Adjust ASCII index to match banner
	}
	return indices
}

// FilterEmptyStrings removes empty strings from a slice of strings.
func FilterEmptyStrings(strings []string) []string {
	var result []string
	for _, str := range strings {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// CalculateFileHash computes the SHA256 hash of a given file.
func CalculateFileHash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to calculate hash: %v", err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// ValidateFileHash validates the file hash against the expected hash.
func ValidateFileHash(filepath, expectedHash string) error {
	hash, err := CalculateFileHash(filepath)
	if err != nil {
		return err
	}
	if hash != expectedHash {
		return errors.New("file integrity check failed: hash mismatch")
	}
	return nil
}

// ContainsNonASCII checks if a string contains non-ASCII characters.
func ContainsNonASCII(text string) bool {
	for _, r := range text {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}
