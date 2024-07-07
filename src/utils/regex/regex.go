package regex

import (
	"regexp"
	"strings"
	"unicode"
)

// Checks if string contains a uppercase letter
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// If string is uppercase, separate string to hyphens
func StringToHyphen(word string) string {
	if IsUpper(word) {
		m1 := regexp.MustCompile(`([a-z])([A-Z])`)

		return strings.ToLower(m1.ReplaceAllString(word, "$1-$2"))
	}

	return strings.ToLower(word)
}

func ToKebabCase(input string) string {
	// Convert uppercase letters to lowercase
	lowercase := strings.ToLower(input)
	// Replace underscores (_) with hyphens (-)
	result := strings.ReplaceAll(lowercase, "_", "-")
	return result
}

// Format snake_case to camelCase (eg. my_variable_name -> myVariableName)
func FormatSnakeCaseToCamelCase(input string) string {
	// Split by underscore
	parts := strings.Split(input, "_")

	// Process each part
	var formattedParts []string
	for i, part := range parts {
		if i == 0 {
			// Leave the first part as is
			formattedParts = append(formattedParts, part)
		} else {
			// Convert subsequent parts to lowercase and uppercase the first letter
			if len(part) > 0 {
				formattedPart := strings.Title(strings.ToLower(part))
				formattedParts = append(formattedParts, formattedPart)
			}
		}
	}

	// Join parts with no underscore
	result := strings.Join(formattedParts, "")

	return result
}
