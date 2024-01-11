package main

import (
	// "fmt"
	// "math"
	b "aoc/utils"
	"fmt"
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
	for _, currentSpring := range s.SpringStatus[0:1] {
		var springStructure = currentSpring.SpringSequence

		var checksums = currentSpring.IntChecksums()
		log.Printf("Spring Line %s %v", springStructure, checksums)
		var matches = structureParser.FindAllStringSubmatch(springStructure, -1)
		var matchIndices = structureParser.FindAllStringSubmatchIndex(springStructure, -1)

		log.Printf("Matches: %v", matches)
		log.Printf("Match indices:%v", matchIndices)
		var matchesLength = b.Map(matches, (func(aMatch []string) int { return len(aMatch[0]) }))
		log.Printf("Match length: %v", matchesLength)

		var partToWorkWith = springStructure
		var remainingChecksumArray = checksums
		log.Printf("partToWorkWith: %s", partToWorkWith)
		log.Printf("remainingChecksumArray: %v", remainingChecksumArray)
		var matchingCombinations, allCombinations = generateValidCombinations(partToWorkWith, remainingChecksumArray)
		log.Printf("No of combinations %d", matchingCombinations)
		log.Printf("%v", allCombinations)
		totalArrangements += matchingCombinations

		var intersections = []rune(allCombinations[0])
		for index, _ := range allCombinations {
			for currentIndex, currentRune := range allCombinations[index] {
				if intersections[currentIndex] != currentRune {
					intersections[currentIndex] = '?'
				}
			}
		}
		log.Printf("%s", string(intersections))
	}

	fmt.Printf("total arrangements %d\n", totalArrangements)

}

func generateValidCombinations(partToWorkWith string, remainingChecksumArray []int) (int, []string) {
	// var fillerMatches = placeholderParser.FindAllString(partToWorkWith, -1)
	var fillerLocations = placeholderParser.FindAllStringSubmatchIndex(partToWorkWith, -1)

	var allCombinations = make([]string, 0)
	// log.Printf("%v, %v", fillerMatches, fillerLocations)
	var alreadyFilled = strings.Count(partToWorkWith, "#")

	var combinationN = len(fillerLocations)
	var combinationR = b.Sum(remainingChecksumArray) - alreadyFilled
	log.Printf("Generating combinations %dC%d", combinationN, combinationR)
	var combinations = combin.Combinations(combinationN, combinationR)
	// log.Printf("%v", combinations)

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
			allCombinations = append(allCombinations, translatedString)
			// distanceCache[translatedString] = matchesLength
		}
		// log.Printf("%s : %v, %v, %v", translatedString, areChecksumsEqual, matchesLength, remainingChecksumArray)
	}
	return matchingCombinations, allCombinations
}

var structurePartParser = regexp.MustCompile("[.?]?[#]+[.?]?")

func (s *SpringField) EvaluatePart2() {

	var totalArrangements = 0
	for _, currentSpring := range s.SpringStatus[0:1] {
		var springStructure = currentSpring.SpringSequence

		var checksums = currentSpring.IntChecksums()

		var endClamp = b.Last(checksums)
		var matches = structurePartParser.FindAllStringSubmatch(springStructure, -1)

		var matchesLengths = b.Map(matches, func(s []string) int { return len(s[0]) })
		var fillerLocations = structurePartParser.FindAllStringSubmatchIndex(springStructure, -1)

		var lastMatchIndices = b.Last(fillerLocations)
		log.Printf("%v %v %v %v %s", endClamp, matchesLengths, lastMatchIndices, matches, springStructure)

		//Try to see if we are missing the start only
		var lastMatch = b.Last(matches)[0]

		var matchingCombinations, choices = generateValidCombinations(lastMatch, []int{endClamp})
		log.Printf("Matching combinations  %d %v", matchingCombinations, choices)

		var addedPlaceholder = "?"
		if lastMatchIndices[1] == len(springStructure) {
			addedPlaceholder = "." //Since we are on the edge of the match
		}

		var addedString = addedPlaceholder + springStructure

		matchingCombinations, _ = generateValidCombinations(addedString, checksums)
		log.Printf("No of combinations for string %s is %d ", addedString, matchingCombinations)

		// for i := 0; i < 5; i++ {
		// 	fmt.Fprintf(&builder, "%s?", springStructure)
		// 	newChecksums = append(newChecksums, checksums...)
		// }
		//
		// var expandedSpring = builder.String()
		// log.Printf("%s %v", expandedSpring, newChecksums)
		//
		// var matches = structureParser.FindAllStringSubmatch(expandedSpring, -1)
		// var matchesLengths = b.Map(matches, func(s []string) int { return len(s[0]) })
		//
		// log.Printf("%v %v", matches, matchesLengths)
		//

		// 	var variations = 0
		// 	var partToWorkWith = springStructure
		// 	var remainingChecksumArray = checksums
		// 	log.Printf("partToWorkWith: %s", partToWorkWith)
		// 	log.Printf("remainingChecksumArray: %v", remainingChecksumArray)
		// 	var matchingCombinations = generateValidCombinations(partToWorkWith, remainingChecksumArray)
		// 	log.Printf("No of combinations %d", matchingCombinations)
		// 	variations = matchingCombinations
		//
		// 	var lastRune = partToWorkWith[len(partToWorkWith)-1]
		// 	var prefix = "?"
		//
		// 	if lastRune == '#' {
		// 		prefix = "."
		// 	}
		//
		// 	var nextString = prefix + partToWorkWith
		// 	// matchingCombinations = generateValidCombinations(nextString, checksums)
		// 	//?###????????
		// 	matchingCombinations = generateValidCombinations(".???????", []int{2, 1})
		// 	log.Printf("No of combinations %d %s %v", matchingCombinations, nextString, checksums)
		// 	variations *= int(math.Pow(float64(matchingCombinations), 4))
		//
		// 	log.Printf("total combinations : %d", variations)
		//
		// 	totalArrangements += variations
	}
	log.Printf("%d", totalArrangements)

}

//TODO: Write a langrange extrapolator value on D3
