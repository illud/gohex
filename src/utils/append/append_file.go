package append

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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
	content, err := os.ReadFile(filename)
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
	err = os.WriteFile(filename, []byte(contentStr), 0644)
	if err != nil {
		return err
	}

	return nil
}

type Tracker struct {
	Modules []struct {
		ModuleName   string `json:"moduleName"`
		EndpointName string `json:"endpointName"`
	} `json:"modules"`
}

func AddDataToTrackerFile(moduleName string, endpointName string) error {
	// Read JSON file
	filePath := "tracker.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal JSON into struct
	var tr Tracker
	if err := json.Unmarshal(data, &tr); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Add new data to the array
	newData := struct {
		ModuleName   string `json:"moduleName"`
		EndpointName string `json:"endpointName"`
	}{
		ModuleName:   moduleName,
		EndpointName: endpointName,
	}
	tr.Modules = append(tr.Modules, newData)

	// Marshal struct back to JSON
	newDataJSON, err := json.MarshalIndent(tr, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Write JSON to file
	err = os.WriteFile(filePath, newDataJSON, os.ModePerm)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	return nil
}

func ReadTrackerFile() Tracker {
	// Read JSON file
	filePath := "tracker.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal JSON into struct
	var tr Tracker
	if err := json.Unmarshal(data, &tr); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return tr
}
