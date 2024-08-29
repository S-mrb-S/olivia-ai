package util

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	bytes := ReadFileContent("res/test/test.txt")

	if string(bytes) != "test" {
		t.Errorf("file.ReadFile() failed.")
	}
}
