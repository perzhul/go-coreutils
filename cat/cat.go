package main

import (
	"fmt"
	"os"
)

func Cat(filePath string) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("could not read file: %s", filePath)
	}
	fmt.Println(string(bytes))

}
