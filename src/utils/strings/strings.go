package strings

import (
	"strings"
	"unicode"
)

func GetFirstCharacterOfString(str string) string {
	return str[0:1]
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

// remove - from string (eg. my-variable-name -> MyVariableName)
func DashToCamel(input string) string {
	words := strings.Split(input, "-")
	for i := 0; i < len(words); i++ {
		words[i] = CapitalizeFirstLetter(words[i])
	}
	return strings.Join(words, "")
}

func CapitalizeFirstLetter(word string) string {
	if len(word) == 0 {
		return ""
	}
	r := []rune(word)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// Format hyphen to camelCase (eg. my-variable-name -> myVariableName)
func FormatHyphenToCamelCase(input string) string {
	// Split by underscore
	parts := strings.Split(input, "-")

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
