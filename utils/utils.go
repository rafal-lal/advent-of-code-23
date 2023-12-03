package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
)

func LoadPuzzleInput(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("error reading file: %s", filePath))
	}

	byteLines := bytes.Split(content, []byte("\n"))
	stringLines := make([]string, len(byteLines))
	for i, val := range byteLines {
		stringLines[i] = string(val)
	}

	return stringLines
}

func SumSlice[T constraints.Ordered](slice []T) T {
	var sum T
	for _, val := range slice {
		sum += val
	}

	return sum
}

func StringInSlice(str string, stringSlice []string) bool {
	for _, val := range stringSlice {
		if val == str {
			return true
		}
	}

	return false
}
