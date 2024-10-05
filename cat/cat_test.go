package main

import (
	"io"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	expectedText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam bibendum, augue et imperdiet tristique, orci enim commodo arcu, elementum pellentesque. \n"

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("can't get working directory")
	}

	filePath := pwd + "/test-file.txt"

	r, w, _ := os.Pipe()

	os.Stdout = w

	Cat(filePath)

	w.Close()

	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("io.ReadAll failed to read the file: %s", filePath)
	}
	result := string(output)

	if len(result) != len(expectedText) {
		t.Fatalf("expected text length to be %d, got %d",
			len(expectedText), len(result))
	}

	if result == expectedText {
		t.Fatalf("expected output of cat function: %s got: %s",
			expectedText, string(output))
	}

}
