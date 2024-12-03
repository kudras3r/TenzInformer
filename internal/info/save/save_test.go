package save

import (
	"os"
	"testing"
)

func TestJSON(t *testing.T) {
	testData := []byte(`{"a": "b"}`)

	tmpFile, err := os.CreateTemp("", "test_*.json")
	if err != nil {
		t.Fatalf("failed to create tmp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = JSON(testData, tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to save test data: %v", err)
	}

	fileData, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read tmp file: %v", err)
	}

	if string(fileData) != string(testData) {
		t.Error("cannot save data")
	}
}
