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
	// Iterate over the files
	// if strings.HasSuffix(file.Name(), ".controller.go") {
	// 	// Get the full path of the file
	// 	// fullPath := filepath.Join(folderPath, file.Name())
	// 	fileName = file.Name()
	// 	// // Print the full path
	// 	// fmt.Println(fullPath)
	// } else {
	// 	return nil, fmt.Errorf("File not found")
	// }
	for _, file := range files {
		// Check if the file has the desired extension
		fileName = file.Name()
		return &fileName, nil
	}

	return nil, errors.New("File not found")
}
