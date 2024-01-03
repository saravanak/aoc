package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000

func Parse(contents string) *OasisReadings {
	var basicParser = participle.MustBuild[OasisReadings](
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

func (s *OasisReadings) Evaluate() {
	s._eval(
		true,
		func(prevLevel int, currentEstimate int) int {
			return prevLevel + currentEstimate
		})
}

func (s *OasisReadings) EvaluatePart2() {
	s._eval(
		false,
		func(prevLevel int, currentEstimate int) int {
			return prevLevel - currentEstimate
		})
}
func (s *OasisReadings) _eval(
	part1 bool,
	diffStackFinder func(prevLevel int, currentEstimate int) int) {
	var sumOfEstimates = 0
	for _, sequenceCommand := range s.Readings {
		var currentSequence = *sequenceCommand.Reading

		var allZeros = false

		var iterationCount = 0
		var diffStack = make([]int, 0)

		for !allZeros {
			var nextSequence = []int{}
			if iterationCount == MAX_ITERS {
				log.Printf("MAX iterationCount reached.")
				break
			}
			if !part1 {
				diffStack = append(diffStack, currentSequence[0])
			}
			var nonZeroIndex = -1
			for i := 1; i < len(currentSequence); i++ {
				var difference = currentSequence[i] - currentSequence[i-1]

				if difference != 0 {
					nonZeroIndex = i
				}
				nextSequence = append(nextSequence, difference)
				if i == len(currentSequence)-1 && part1 {
					diffStack = append(diffStack, currentSequence[i])
				}
			}
			currentSequence = nextSequence
			// log.Printf("%+v", currentSequence)
			allZeros = nonZeroIndex == -1
			iterationCount++
		}

		if !allZeros {
			panic(fmt.Sprintf("Unable to reach end condition for Sequence :%+v", sequenceCommand.Reading))
		}

		// log.Printf("%+v", diffStack)
		var currentEstimate = 0
		for i := len(diffStack) - 1; i >= 0; i-- {
			currentEstimate = diffStackFinder(diffStack[i], currentEstimate)
		}

		log.Printf("Next Estimate value: %d", currentEstimate)
		sumOfEstimates += currentEstimate
	}

	log.Printf("Part 1 solution %d", sumOfEstimates)
}

//TODO: Write a langrange extrapolator value on D3
