package main

import (
	"log"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000

func Parse(contents string) *ProgramText {
	var basicParser = participle.MustBuild[ProgramText](
		participle.Lexer(basicLexer),
		participle.UseLookahead(2),
	)

	ast, err := basicParser.ParseString("", contents)
	if err != nil {
		log.Println("Parse String error", err)
		return nil
	}
	return ast
}

func (s *ProgramText) Evaluate() {

	// log.Printf("%v", s.Commands)

	var sum = 0
	for _, command := range *s.Commands {
		var currentHash = calcWeights(command.Word)
		log.Printf("'%s' :%d", command.Word, currentHash)
		sum += currentHash
	}
	log.Printf("%d", sum)
}

var weights = make(map[string]int)

func calcWeights(input string) int {
	//   Determine the ASCII code for the current character of the string.
	// Increase the current value by the ASCII code you just determined.
	// Set the current value to itself multiplied by 17.
	// Set the current value to the remainder of dividing itself by 256.

	var result = 0

	// log.Printf("%v", weights)
	var foundCachedValue = false
	var currentSearchIndex = len(input)

	for !foundCachedValue && currentSearchIndex >= 0 {
		var searchString = input[0:currentSearchIndex]
		result = weights[searchString]
		foundCachedValue = result > 0
		// log.Printf("%v %d %s", foundCachedValue, result, searchString)
		if !foundCachedValue {
			currentSearchIndex--
		}
	}

	var startIndex = 0
	if foundCachedValue {
		// log.Printf("Using cache %s %d %d", input[0:currentSearchIndex], result, currentSearchIndex)
		startIndex = currentSearchIndex
	}
	for i := startIndex; i < len(input); i++ {
		var asciiValue = int(input[i])
		// log.Printf("%d %s", asciiValue, input[startIndex:i+1])
		result += asciiValue
		result *= 17
		result %= 256
		weights[input[0:i+1]] = result
	}

	return result
}

func (s *ProgramText) EvaluatePart2() {

}
