package main

import (
	// "fmt"
	// "fmt"
	b "aoc/utils"
	"log"

	"github.com/alecthomas/participle/v2"
	"regexp"
)

var MAX_ITERS = 10000

func Parse(contents string) *SpringField {
	var basicParser = participle.MustBuild[SpringField](
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

func (s *SpringField) Evaluate() {
	var structureParser = regexp.MustCompile("[#]+")
	for _, currentSpring := range s.SpringStatus {
		var springStructure = currentSpring.SpringSequence
		log.Printf("%s", springStructure)

		log.Printf("%v", b.Map(currentSpring.Checksum, (func(checksum Checksum) int { return checksum.SingleChecksum })))
		log.Printf("%v", structureParser.FindAllStringSubmatch(springStructure, -1))
	}

}

func (s *SpringField) EvaluatePart2() {

}

//TODO: Write a langrange extrapolator value on D3
