package test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsMainFileExists(t *testing.T) {
	// Check main.go file
	_, err := os.Stat("./../../main.go")
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatal("File main.go not found in the root directory")
		} else {
			t.Fatalf("An error occurred while checking the main.go file: %v", err)
		}
	}
}

func TestIsGoModFileExists(t *testing.T) {
	// Check main.go file
	_, err := os.Stat("./../../go.mod")
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatal("File go.mod not found in the root directory")
		} else {
			t.Fatalf("An error occurred while checking the go.mod file: %v", err)
		}
	}
}

func TestIsDatabaseFileExists(t *testing.T) {
	// Check database.go file
	_, err := os.Stat("./../../database/database.go")
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatal("File database.go not found in the root directory")
		} else {
			t.Fatalf("An error occurred while checking the database.go file: %v", err)
		}
	}
}

func TestIsModelFileExists(t *testing.T) {
	// Specify the folder to check
	folder := "./../../model/"

	// Get a list of files in the folder
	fileList, err := filepath.Glob(filepath.Join(folder, "*.go"))
	if err != nil {
		t.Fatalf("An error occurred while getting the file list: %v", err)
	}

	// Check if any model file exists in the folder
	if len(fileList) == 0 {
		t.Fatalf("No model files found in the folder: %s", folder)
	}

	// Print the list of .go files found
	t.Logf("Found %d .go files in the folder: %s", len(fileList), folder)
	for _, file := range fileList {
		t.Log(file)
	}
}

func TestIsControllerFileExists(t *testing.T) {
	// Specify the folder to check
	folder := "./../../controller/"

	// Get a list of files in the folder
	fileList, err := filepath.Glob(filepath.Join(folder, "*.go"))
	if err != nil {
		t.Fatalf("An error occurred while getting the file list: %v", err)
	}

	// Check if any controller file exists in the folder
	if len(fileList) == 0 {
		t.Fatalf("No controller files found in the folder: %s", folder)
	}

	// Print the list of .go files found
	t.Logf("Found %d .go files in the folder: %s", len(fileList), folder)
	for _, file := range fileList {
		t.Log(file)
	}
}
