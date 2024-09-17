package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var splitChars = []string{"_", "-", "."}

func isSplitChar(char rune) bool {
	for _, splitChar := range splitChars {
		if char == rune(splitChar[0]) {
			return true
		}
	}
	return false
}

func toPascalCase(s string) string {
	var splitChar = ""
	for _, char := range splitChars {
		if strings.Contains(s, char) {
			splitChar = char
			break
		}
	}

	if splitChar != "" {
		words := strings.Split(s, splitChar)
		for i := 0; i < len(words); i++ {
			words[i] = cases.Title(language.English).String(words[i])
		}
		return strings.Join(words, "")
	}

	return cases.Title(language.English).String(s)
}

func toScreamingSnakeCase(s string) string {
	var result strings.Builder
	var prevChar rune
	for i, char := range s {
		if i > 0 {
			if unicode.IsUpper(char) && !unicode.IsUpper(prevChar) && !isSplitChar(prevChar) {
				result.WriteRune('_')
			} else if isSplitChar(char) && !isSplitChar(prevChar) {
				result.WriteRune('_')
				prevChar = char
				continue
			}
		}

		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			result.WriteRune(unicode.ToUpper(char))
		} else if !isSplitChar(char) {
			result.WriteRune('_')
		}

		prevChar = char
	}

	return result.String()
}
