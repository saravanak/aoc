package main

import (
	"4d63.com/strrev"
	b "aoc/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {

	//input_file := "./data/01/example.txt"
	input_file := "./data/01/day-01.txt"
	fileLines := b.ReadFileAsArray(input_file)

	match, _ := regexp.MatchString("(\\d)", fileLines[0])
	log.Println(match)

	matcherRegexp, _ := regexp.Compile("(\\d)")

	sum := 0

	for _, calibWord := range fileLines {
		log.Println(calibWord)
		match_slices := matcherRegexp.FindAllString(calibWord, -1)
		two_digits := match_slices[0] + match_slices[len(match_slices)-1]
		as_number, _ := strconv.Atoi(two_digits)

		sum += as_number
	}

	fmt.Println(sum)
}

func part2() {
	// input_file := "./data/01/example-2.txt"
	input_file := "./data/01/day-01.txt"
	fileLines := b.ReadFileAsArray(input_file)

	digit_names := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var reversed_names []string

	for _, digit_word := range digit_names {
		reversed_names = append(reversed_names, strrev.Reverse(digit_word))
	}

	reversed_regex_source_string := fmt.Sprintf("(%s|\\d)", strings.Join(reversed_names, "|"))
	normal_regex_source_string := fmt.Sprintf("(%s|\\d)", strings.Join(digit_names[:], "|"))
	reversedRegexp := regexp.MustCompile(reversed_regex_source_string)
	normalRegexp := regexp.MustCompile(normal_regex_source_string)

	sum := 0

	for _, calibWord := range fileLines {
		log.Println(calibWord)
		startMatch := normalRegexp.FindAllString(calibWord, -1)
		endMatch := reversedRegexp.FindAllString(strrev.Reverse(calibWord), -1)

		// Reverse the matches to get the original string back.
		var endMatches []string
		for _, endMatchItem := range endMatch {
			endMatches = append(endMatches, strrev.Reverse(endMatchItem))
		}

		number_slices := as_number_slice(startMatch)
		end_slices := as_number_slice(endMatches)

		two_digits := number_slices[0] + end_slices[0]
		log.Println(number_slices, two_digits)
		as_number, _ := strconv.Atoi(two_digits)

		sum += as_number
	}

	fmt.Println(sum)

}

func as_number_slice(matches []string) []string {
	var number_slice []string
	for _, v := range matches {
		switch v {
		case "one":
			number_slice = append(number_slice, "1")
		case "two":
			number_slice = append(number_slice, "2")
		case "three":
			number_slice = append(number_slice, "3")
		case "four":
			number_slice = append(number_slice, "4")
		case "five":
			number_slice = append(number_slice, "5")
		case "six":
			number_slice = append(number_slice, "6")
		case "seven":
			number_slice = append(number_slice, "7")
		case "eight":
			number_slice = append(number_slice, "8")
		case "nine":
			number_slice = append(number_slice, "9")
		default:
			number_slice = append(number_slice, v)
		}
	}
	return number_slice

}
func main() {

	if os.Args[1] == "part2" {
		fmt.Println("running day02/part 02!!!")
		part2()
	} else {
		fmt.Println("running day02/part 01!!!")
		part1()
	}
}
