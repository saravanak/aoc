package main

import (
	"fmt"
	"regexp"
	// "strings"
	"testing"
)

func TestRegexMatches(t *testing.T) {

	input_text := "...2323...454..#..$..*..<..>..+..!"
	parsePartsAndSymbols(input_text)
	t.Errorf("I am fail")
}
