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

var structureParser = regexp.MustCompile("[#]+")
var placeholderParser = regexp.MustCompile("[?]")
var distanceCache = make(map[string][]int)

func (s *SpringField) Evaluate() {
	var totalArrangements = 0
	for _, currentSpring := range s.SpringStatus {
		var springStructure = currentSpring.SpringSequence

		var checksums = currentSpring.IntChecksums()
		log.Printf("Spring Line %s %v", springStructure, checksums)
		var matches = structureParser.FindAllStringSubmatch(springStructure, -1)
		var matchIndices = structureParser.FindAllStringSubmatchIndex(springStructure, -1)

		log.Printf("Matches: %v", matches)
		log.Printf("Match indices:%v", matchIndices)
		var matchesLength = b.Map(matches, (func(aMatch []string) int { return len(aMatch[0]) }))
		log.Printf("Match length: %v", matchesLength)

		// distanceCache[springStructure[b.Last(matchIndices)[0]:]] = []int{b.Last(checksums)}
		// var matchIndices = structureParser.FindAllStringSubmatchIndex(translatedString, -1)
		// var leftEndClamp, rightEndClamp = b.Clamp(true, checksums, matchesLength)

		// if leftEndClamp >= 0 {
		// 	log.Printf("We've got a end clamp for %s: Left: %d , Right:%d", springStructure, leftEndClamp, rightEndClamp)
		// 	var partToWorkWith = springStructure[0:matchIndices[rightEndClamp][0]]
		// 	var remainingChecksumArray = checksums[0:leftEndClamp]
		// 	log.Printf("partToWorkWith: %s", partToWorkWith)
		// 	log.Printf("remainingChecksumArray: %v", remainingChecksumArray)
		// 	var matchingCombinations = generateValidCombinations(partToWorkWith, remainingChecksumArray)
		// 	totalArrangements += matchingCombinations
		// 	log.Printf("No of combinations %d", matchingCombinations)
		// 	continue
		// }
		// var leftStartClamp, rightStartClamp = b.Clamp(false, checksums, matchesLength)
		//
		// if leftStartClamp >= 0 {
		// 	log.Printf("We've got a start clamp for %s: Left: %d , Right:%d", springStructure, leftStartClamp, rightStartClamp)
		// 	var partToWorkWith = springStructure[matchIndices[rightStartClamp][1]+1:]
		// 	var remainingChecksumArray = checksums[leftStartClamp+1:]
		// 	log.Printf("partToWorkWith: %s", partToWorkWith)
		// 	log.Printf("remainingChecksumArray: %v", remainingChecksumArray)
		// 	var matchingCombinations = generateValidCombinations(partToWorkWith, remainingChecksumArray)
		// 	totalArrangements += matchingCombinations
		// 	log.Printf("No of combinations %d", matchingCombinations)
		// 	continue
		// }
		var partToWorkWith = springStructure
		var remainingChecksumArray = checksums
		log.Printf("partToWorkWith: %s", partToWorkWith)
		log.Printf("remainingChecksumArray: %v", remainingChecksumArray)
		var matchingCombinations = generateValidCombinations(partToWorkWith, remainingChecksumArray)
		log.Printf("No of combinations %d", matchingCombinations)
		totalArrangements += matchingCombinations
		// log.Printf("%v", distanceCache)
	}

	log.Printf("total arrangements %d", totalArrangements)

}

func generateValidCombinations(partToWorkWith string, remainingChecksumArray []int) int {
	var fillerMatches = placeholderParser.FindAllString(partToWorkWith, -1)
	var fillerLocations = placeholderParser.FindAllStringSubmatchIndex(partToWorkWith, -1)

	log.Printf("%v, %v", fillerMatches, fillerLocations)
	var alreadyFilled = strings.Count(partToWorkWith, "#")

	var combinationN = len(fillerLocations)
	var combinationR = b.Sum(remainingChecksumArray) - alreadyFilled
	log.Printf("Generating combinations %dC%d", combinationN, combinationR)
	var combinations = combin.Combinations(combinationN, combinationR)
	log.Printf("%v", combinations)

	var matchingCombinations = 0
	for _, currentCombination := range combinations {
		var newStringRune = []rune(partToWorkWith)

		for _, combinationPicker := range currentCombination {
			newStringRune[fillerLocations[combinationPicker][0]] = rune('#')
		}

		var newString = string(newStringRune)
		var translatedString = strings.ReplaceAll(newString, "?", ".")
		var matches = structureParser.FindAllStringSubmatch(translatedString, -1)
		var matchesLength = b.Map(matches, (func(aMatch []string) int { return len(aMatch[0]) }))
		var areChecksumsEqual = slices.Equal(matchesLength, remainingChecksumArray)

		if areChecksumsEqual {
			matchingCombinations++
			// distanceCache[translatedString] = matchesLength
		}
		log.Printf("%s : %v, %v, %v", translatedString, areChecksumsEqual, matchesLength, remainingChecksumArray)
	}
	return matchingCombinations
}

func (s *SpringField) EvaluatePart2() {

}

//TODO: Write a langrange extrapolator value on D3
