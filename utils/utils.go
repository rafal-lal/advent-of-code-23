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

func RemoveOrdered[T any](slice []T, idx int) []T {
	if idx == len(slice)-1 {
		return slice[:idx]
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func RemoveUnordered[T any](slice []T, idx int) []T {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func StringInSlice(str string, stringSlice []string) bool {
	for _, val := range stringSlice {
		if val == str {
			return true
		}
	}

	return false
}

func MinInSlice[T constraints.Ordered](data []T) T {
	var min T
	for i, val := range data {
		if i == 0 {
			min = val
		}
		if val < min {
			min = val
		}
	}

	return min
}
