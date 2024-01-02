package main

import (
	// b "aoc/utils"
	"github.com/alecthomas/participle/v2"
	"log"
)

var MAX_ITERS = 10

func Parse(contents string) *RacingRecords {

	// parser, parse_error := participle.Build[SEED_PLAN]()
	var basicParser = participle.MustBuild[RacingRecords](
		// participle.Lexer(basicLexer),
		participle.UseLookahead(2),
	)

	ast, err := basicParser.ParseString("", contents)
	if err != nil {
		log.Println("Parse String error", err)
		return nil
	}
	return ast
}

func (s *RacingRecords) Evaluate() {

	var times = *s.Commands[0].TimeCommand.RaceTime
	var distances = *s.Commands[1].DistanceCommand.Distances

	var sum = 1
	for raceIndex := 0; raceIndex < len(times); raceIndex++ {
		// for raceIndex := 2; raceIndex < 3; raceIndex++ {
		var high = gradientUp(times[raceIndex], distances[raceIndex])
		var low = gradientDown(times[raceIndex], distances[raceIndex])

		sum *= (high - low + 1)
		log.Printf("%d->%d", low, high)

	}

	log.Printf("Solution is %d", sum)
}

func gradientUp(totalTime int, targetDistance int) int {
	var high = totalTime
	var low = 0
	var prevHigh = high
	var (
		currentDistance     = 0
		nextMsDistance  int = 0
	)
	var stopCondition = false
	var circuitBreaker = 0

	log.Printf("Finding upper bound for %d ms over distance %d", totalTime, targetDistance)

	for !stopCondition {
		circuitBreaker++
		if circuitBreaker > MAX_ITERS || low >= high {
			log.Printf("CIRCUIT BREAKING!!")
			break
		}
		log.Printf("Low: %d; High : %d; prevHigh: %d", low, high, prevHigh)
		currentDistance = distanceCovered(high, totalTime)
		nextMsDistance = distanceCovered(high+1, totalTime)

		log.Printf("currentDistance: %d; nextMsDistance : %d", currentDistance, nextMsDistance)
		if currentDistance > targetDistance && nextMsDistance <= targetDistance {
			stopCondition = true
			break
		}
		if currentDistance < targetDistance {
			prevHigh = high
			high = low + (high-low)/2 - 1
		} else {
			low = high
			high = high + (prevHigh-low)/2
		}
	}
	return high

}

func gradientDown(totalTime int, targetDistance int) int {
	var high = totalTime
	var low = 0
	var prevLow = low
	var (
		currentDistance = 0
		prevMsDistance  = 0
	)
	var stopCondition = false
	var circuitBreaker = 0

	log.Printf("Finding lower bound for %d ms over distance %d", totalTime, targetDistance)

	for !stopCondition {
		circuitBreaker++
		if circuitBreaker > MAX_ITERS || low >= high {
			log.Printf("CIRCUIT BREAKING!!")
			break
		}
		log.Printf("Low: %d; High : %d; currentBound: %d", low, high, prevLow)
		currentDistance = distanceCovered(low, totalTime)
		prevMsDistance = distanceCovered(low-1, totalTime)

		if currentDistance >= targetDistance && prevMsDistance <= targetDistance {
			stopCondition = true
			break
		}
		log.Printf("currentDistance: %d ", currentDistance)
		if currentDistance < targetDistance {
			prevLow = low
			low = low + (high-low)/2
		} else {
			high = low
			low = low + (prevLow-low)/2
		}
	}
	return low

}
func distanceCovered(waitTime int, totalTime int) int {
	var startingSpeed = waitTime
	var timeTravelling = totalTime - waitTime

	return startingSpeed * timeTravelling
}
func (s *RacingRecords) EvaluatePart2() {

}
