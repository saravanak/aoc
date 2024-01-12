package utils

import (
	"bufio"
	"log"
	"os"
	"slices"
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

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Sum(slice []int) int {
	var result = 0
	for _, value := range slice {

		result += value
	}
	return result
}
func IntComparer(a int, b int) int {
	if a > b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Finds the indices of respective arrays starting at which the two arrays are equal till the end.
// Returns index >=0 if found else -1 for each input array in the respective order
func Clamp[T comparable](isEnd bool, lhs []T, rhs []T) (int, int) {
	var strategy = "end"

	if !isEnd {
		strategy = "start"
	}

	log.Printf("%sClamp lhs: %v, rhs: %v", strategy, lhs, rhs)
	var workingLhs = slices.Clone(lhs)
	var workingRhs = slices.Clone(rhs)

	if isEnd {

		slices.Reverse(workingLhs)
		slices.Reverse(workingRhs)
	}

	var endClamp = -1
	for index := 0; index < len(workingLhs); index++ {

		if index >= len(workingRhs) {
			break
		}
		if workingLhs[index] != workingRhs[index] {
			break
		}
		endClamp = index
	}
	if endClamp == -1 {
		return endClamp, endClamp
	}
	if isEnd {
		return len(workingLhs) - endClamp - 1, len(workingRhs) - endClamp - 1
	} else {
		return endClamp, endClamp
	}

}

func Transpose[T any](input [][]T) [][]T {

	var returnVal = make([][]T, len(input[0]))

	for lineIndex, currentLine := range input {
		for columnIndex, currentRune := range currentLine {
			if returnVal[columnIndex] == nil {
				returnVal[columnIndex] = make([]T, len(input))
			}
			returnVal[columnIndex][lineIndex] = currentRune
		}
	}
	log.Printf("%v %d %d", returnVal, len(returnVal), len(returnVal[0]))
	return returnVal
}
