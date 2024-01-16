package main

import (
	b "aoc/utils"
	"bytes"
	"fmt"
	"log"

	"github.com/alecthomas/participle/v2"
	"golang.org/x/exp/slices"
)

// var MAX_ITERS = 1000000000

var MAX_ITERS = 300000

type directionIterators struct {
	outer          []int
	inner          []int
	slideDirection int
	swapLineCols   bool
}

func Parse(contents string) *PipeMap {
	var basicParser = participle.MustBuild[PipeMap](
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

func (s *PipeMap) Evaluate() {

	s.join()

	var rockLocations bytes.Buffer

	for i := 0; i < len(s.RuneMatrix); i++ {
		for j := 0; j < len(s.RuneMatrix[i]); j++ {
			if s.RuneMatrix[i][j] == 'O' {
				fmt.Fprintf(&rockLocations, "(%d:%d),", i, j)
			}
		}
	}

	type rotationMap map[string]string

	var fromState = b.GetSha256(&rockLocations)
	var locationsMap = make(map[string]rotationMap)

	var directions = []string{"n", "w", "s", "e"}

	for _, direction := range directions {
		locationsMap[direction] = make(rotationMap)
	}
	var stateVsIteration = make(map[string]int)
	var loads = make(map[string]int)
	var prevState = ""

	var i = 0
	var cacheStartState = ""
	for ; i < MAX_ITERS*4; i++ {
		var currentDirection = directions[i%4]

		if i%100000 == 0 {
			log.Printf("%d", i)
		}

		var currentLocation = locationsMap[currentDirection]
		if currentLocation[fromState] != "" {
			// log.Printf("Previous occurence %d currentIteration: %d, direction %s, currentIteration %d", stateVsIteration[prevState], stateVsIteration[fromState], currentDirection, i)
			prevState = fromState
			fromState = currentLocation[fromState]
		} else {
			var outputHash string
			stateVsIteration[fromState] = i
			outputHash, _ = slide(s)
			currentLocation[fromState] = outputHash
			if cacheStartState != "" {
				panic("Load go we ")
			}
			prevState = fromState
			fromState = outputHash

			s.RuneMatrix = slices.Clone(b.RotateClockwise(s.RuneMatrix))
		}
		if currentDirection == "e" {
			var load = s.calcLoads()
			// log.Printf("%d", loads[fromState])
			var currentState = prevState + fromState
			if load == 64 {
				log.Printf("Got target load on currentIteration : %d %s", i, currentState)
			}
			loads[currentState] = load
			// log.Printf("After Rotating %s, Cycle: %d ", currentDirection, (i+1)/4)
			// for _, line := range s.RuneMatrix {
			// 	log.Printf("%v ", string(line))
			// }
		}
	}
	log.Printf("%v", loads)

}

func slide(s *PipeMap) (string, int) {
	var sum = 0
	var rockLocations bytes.Buffer
	var lastInsertPositions = make([]int, len(s.RuneMatrix[0]))
	rockLocations.Reset()
	for line := 0; line < len(s.RuneMatrix); line++ {
		for col := 0; col < len(s.RuneMatrix[line]); col++ {
			if s.RuneMatrix[line][col] == '.' {
				if lastInsertPositions[col] == -1 {
					lastInsertPositions[col] = line
				}
			}
			if s.RuneMatrix[line][col] == '#' {
				lastInsertPositions[col] = -1
			}
			if s.RuneMatrix[line][col] == 'O' {
				if lastInsertPositions[col] != -1 {
					var distance = (len(s.RuneMatrix) - lastInsertPositions[col])
					sum += distance
					fmt.Fprintf(&rockLocations, "(%d:%d),", line, lastInsertPositions[col])
					s.RuneMatrix[lastInsertPositions[col]][col], s.RuneMatrix[line][col] = s.RuneMatrix[line][col], s.RuneMatrix[lastInsertPositions[col]][col]
					// log.Printf("Moving rock from [ %d,%d ] to line %d", line, col, lastInsertPositions[col])
					lastInsertPositions[col] += 1
					// log.Printf("%d is the distance for %d %d", distance, line, col)
					// log.Printf("New insert positin for column %d is %d  ", col, lastInsertPositions[col])
				} else {
					var distance = len(s.RuneMatrix) - line
					// log.Printf("%d is the distance for %d %d", distance, line, col)
					sum += distance
					fmt.Fprintf(&rockLocations, "(%d:%d),", line, col)
				}
			}
		}
	}
	return b.GetSha256(&rockLocations), sum

}

func (s *PipeMap) EvaluatePart2() {

}
