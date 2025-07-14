package find

import (
	"errors"
	"os"
)

func FindFile(folderPath string) (*string, error) {
	// Read the contents of the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var fileName string

	for _, file := range files {
		// Check if the file has the desired extension
		fileName = file.Name()
		return &fileName, nil
	}

	return nil, errors.New("File not found")
}
