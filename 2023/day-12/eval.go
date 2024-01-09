package main

import (
	// "fmt"
	// "fmt"
	b "aoc/utils"
	"log"
	"slices"

	"regexp"
	"strings"

	"gonum.org/v1/gonum/stat/combin"

	"github.com/alecthomas/participle/v2"
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

	var fillterParser = regexp.MustCompile("[?]")
	for _, currentSpring := range s.SpringStatus[0:1] {
		var springStructure = currentSpring.SpringSequence
		log.Printf("%s", springStructure)

		var checksums = b.Map(currentSpring.Checksum, (func(checksum Checksum) int { return checksum.SingleChecksum }))
		log.Printf("%v", checksums)
		var matches = structureParser.FindAllStringSubmatch(springStructure, -1)
		var matchIndices = structureParser.FindAllStringSubmatchIndex(springStructure, -1)

		log.Printf("Matches: %v", matches)
		log.Printf("Match indices:%v", matchIndices)
		var matchesLength = b.Map(matches, (func(aMatch []string) int { return len(aMatch[0]) }))
		log.Printf("Match length: %v", matchesLength)

		if b.Last(checksums) == b.Last(matchesLength) {
			log.Printf("We've got a end clamp for %s", springStructure)

			var partToWorkWith = springStructure[0:b.Last(matchIndices)[0]]
			var remainingChecksumArray = checksums[0 : len(checksums)-1]
			log.Printf("%s", partToWorkWith)
			log.Printf("%v", remainingChecksumArray)

			var fillerMatches = fillterParser.FindAllString(partToWorkWith, -1)
			var fillerLocations = fillterParser.FindAllStringSubmatchIndex(partToWorkWith, -1)

			log.Printf("%v, %v", fillerMatches, fillerLocations)

			var combinations = combin.Combinations(len(fillerMatches), len(remainingChecksumArray))
			log.Printf("%v", combinations)

			for _, currentCombination := range combinations {
				var newStringRune = []rune(partToWorkWith)

				for _, combinationPicker := range currentCombination {
					newStringRune[fillerLocations[combinationPicker][0]] = rune('#')
				}

				var newString = string(newStringRune)

				var translatedString = strings.ReplaceAll(newString, "?", ".")

				log.Printf("%s", translatedString)

				var matches = structureParser.FindAllStringSubmatch(translatedString, -1)
				// var matchIndices = structureParser.FindAllStringSubmatchIndex(translatedString, -1)

				var matchesLength = b.Map(matches, (func(aMatch []string) int { return len(aMatch[0]) }))
				var areChecksumsEqual = slices.Equal(matchesLength, remainingChecksumArray)
				log.Printf("%s : %v, %v, %v", translatedString, areChecksumsEqual, matchesLength, remainingChecksumArray)
			}
		}
	}

}

func (s *SpringField) EvaluatePart2() {

}

//TODO: Write a langrange extrapolator value on D3
