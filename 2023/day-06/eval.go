package main

import (
	// b "aoc/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000

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

func (s *RacingRecords) EvaluatePart2() {

	var times = *s.Commands[0].TimeCommand.RaceTime
	var distances = *s.Commands[1].DistanceCommand.Distances

	var timeString = ""
	var distanceString = ""
	for raceIndex := 0; raceIndex < len(times); raceIndex++ {
		timeString += fmt.Sprintf("%d", times[raceIndex])
		distanceString += fmt.Sprintf("%d", distances[raceIndex])
	}
	// for raceIndex := 2; raceIndex < 3; raceIndex++ {
	var fullTime, _ = strconv.Atoi(timeString)
	var fullDistance, _ = strconv.Atoi(distanceString)
	var high = gradientUp(fullTime, fullDistance)
	var low = gradientDown(fullTime, fullDistance)

	var sum = (high - low + 1)

	log.Printf("Solution is %d", sum)
}

func gradientUp(totalTime int, targetDistance int) int {
	var high = totalTime
	var low = 0
	var prevHigh = high
	var direction = -1 // +1 for right and -1 for left.
	var (
		currentDistance     = 0
		nextMsDistance  int = 0
	)
	var stopCondition = false
	var circuitBreaker = 0

	log.Printf("Finding upper bound for %d ms over distance %d", totalTime, targetDistance)

	for !stopCondition {
		circuitBreaker++
		if circuitBreaker > MAX_ITERS {
			log.Printf("CIRCUIT BREAKING!!")
			break
		}

		if direction > 0 {
			low = high
			high = high + (totalTime-high)/2
		} else {
			high = high - (high-low)/2
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
			direction = -1
		} else {
			direction = 1
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
	var direction = +1 // +1 for right and -1 for left.
	var stopCondition = false
	var circuitBreaker = 0

	log.Printf("Finding lower bound for %d ms over distance %d", totalTime, targetDistance)

	for !stopCondition {
		circuitBreaker++
		if circuitBreaker > MAX_ITERS {
			log.Printf("CIRCUIT BREAKING!!")
			break
		}

		log.Printf("Low: %d; High : %d; currentBound: %d, direction : %d", low, high, prevLow, direction)
		if direction > 0 {
			low = low + (high-low)/2
		} else {
			prevLow = low
			low = (high - low) / 2
			high = prevLow
		}
		log.Printf("Low: %d; High : %d", low, high)
		currentDistance = distanceCovered(low, totalTime)
		prevMsDistance = distanceCovered(low-1, totalTime)

		if currentDistance > targetDistance && prevMsDistance <= targetDistance {
			stopCondition = true
			break
		}
		log.Printf("currentDistance: %d ", currentDistance)
		if currentDistance < targetDistance {
			direction = 1
		} else {
			direction = -1
		}
	}
	return low

}
func distanceCovered(waitTime int, totalTime int) int {
	var startingSpeed = waitTime
	var timeTravelling = totalTime - waitTime

	return startingSpeed * timeTravelling
}
