package util

import (
	"os"
)

// ReadFileContent returns the bytes of a file searched in the path and beyond it
/*
Renamed `ReadFile` to `ReadFileContent` for clarity.
Renamed `path` to `filePath` to specify it's the file path.
Renamed `bytes` to `fileBytes` to indicate it's the file content.
Renamed `err` to `readError` for clarity on the type of error.
*/
func ReadFileContent(filePath string) (fileBytes []byte) {
	fileBytes, readError := os.ReadFile(filePath)
	if readError != nil {
		fileBytes, readError = os.ReadFile("../" + filePath)
	}

	if readError != nil {
		panic(readError)
	}

	return fileBytes
}
