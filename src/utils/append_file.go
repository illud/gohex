package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func AppendDataToFile(filePath string, data string) error {
	// Open the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write the data to the end of the file
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func ReplaceLastCharacter(filename string, oldChar, newChar string) error {
	// Read the entire file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert content to string for easier manipulation
	contentStr := string(content)

	// Find the last occurrence of the old character
	lastIndex := strings.LastIndex(contentStr, oldChar)
	if lastIndex == -1 {
		return fmt.Errorf("'%s' not found in the file", oldChar)
	}

	// Replace the last occurrence with the new character
	contentStr = contentStr[:lastIndex] + newChar + contentStr[lastIndex+1:]

	// Write the modified content back to the file
	err = ioutil.WriteFile(filename, []byte(contentStr), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Last character replaced successfully.")
	return nil
}
