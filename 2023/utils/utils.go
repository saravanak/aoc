package utils

import (
	"bufio"
	"os"
)

func ReadFileAsArray(name string) []string {
	readFile, err := os.Open(name)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Filter[T any](slice []T, comp func(T) bool) []T {

	returnVal := make([]T, 0)
	for _, item := range slice {
		predicateValue := comp(item)
		if predicateValue {
			returnVal = append(returnVal, item)
		}
	}
	return returnVal

}

func Map[T any, R any](slice []T, mapper func(T) R) []R {

	returnVal := make([]R, 0)
	for _, item := range slice {
		predicateValue := mapper(item)
		returnVal = append(returnVal, predicateValue)
	}
	return returnVal

}
func IntComparer(a int, b int) int {
	if a > b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
