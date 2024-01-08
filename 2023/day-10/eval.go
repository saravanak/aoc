package main

import (
	// "fmt"
	"github.com/alecthomas/participle/v2"
	"log"
	"strconv"
)

var MAX_ITERS = 10000

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

var collectedPoints = make(map[int]*Fitting, 0)

func (s *PipeMap) Evaluate() {
	log.Printf("%+v", s.Line[1].Fitting)

	var start2Fitting *Fitting
	var generalFitting *Fitting
	for _, line := range s.Line {

		for index := range line.Fitting {
			generalFitting = &line.Fitting[index]

			if *generalFitting.Type == "S" {
				start2Fitting = &line.Fitting[index]
				log.Printf("Got start node at %d", start2Fitting.Pos.Column)
			}

			switch *generalFitting.Type {
			case ".":
				generalFitting.FittingBehaviour = NonePipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "-":
				generalFitting.FittingBehaviour = &HorizontalPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "|":
				generalFitting.FittingBehaviour = VertPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "L":
				generalFitting.FittingBehaviour = LPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "J":
				generalFitting.FittingBehaviour = JPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "F":
				generalFitting.FittingBehaviour = FPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "7":
				generalFitting.FittingBehaviour = SevenPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
			case "S":
				generalFitting.FittingBehaviour = HorizontalPipeFitting{BaseFitting: BaseFitting{Pos: generalFitting.Pos}}
				log.Printf("adj %+v", generalFitting.FittingBehaviour.getAdjacentFittings(s))
			}
		}
	}

	var directionFlags int64 = 0 // BTLR -> 1 if the S cell has a join on that direction
	var adjacentCellsToSource = start2Fitting.FittingBehaviour.getAdjacentFittings(s)
	// log.Printf("%s", *adjacentCellsToSource.left.Type)
	// log.Printf("%s", *adjacentCellsToSource.right.Type)
	// log.Printf("%s", *adjacentCellsToSource.top.Type)
	// log.Printf("%s", *adjacentCellsToSource.bottom.Type)

	if adjacentCellsToSource.left != nil && adjacentCellsToSource.left.FittingBehaviour.opensRight() {
		directionFlags |= 2
	}

	if adjacentCellsToSource.right != nil && adjacentCellsToSource.right.FittingBehaviour.opensLeft() {
		directionFlags |= 1
	}
	if adjacentCellsToSource.top != nil && adjacentCellsToSource.top.FittingBehaviour.opensBottom() {
		directionFlags |= 4
	}
	if adjacentCellsToSource.bottom != nil && adjacentCellsToSource.bottom.FittingBehaviour.opensTop() {
		directionFlags |= 8
	}

	var asBinary = strconv.FormatInt(directionFlags, 2)
	log.Printf("%s", asBinary)

	switch asBinary { //BTLR
	case "1100": //TB
		*start2Fitting.Type = "|"
		start2Fitting.FittingBehaviour = VertPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	case "1010": //BL
		*start2Fitting.Type = "7"
		start2Fitting.FittingBehaviour = SevenPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	case "1001": //BR
		*start2Fitting.Type = "F"
		start2Fitting.FittingBehaviour = FPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	case "110": //TL
		*start2Fitting.Type = "J"
		start2Fitting.FittingBehaviour = JPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	case "101": //TR
		*start2Fitting.Type = "L"
		start2Fitting.FittingBehaviour = LPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	case "11": //LR
		*start2Fitting.Type = "-"
		start2Fitting.FittingBehaviour = HorizontalPipeFitting{BaseFitting: BaseFitting{start2Fitting.Pos}}
	}

	log.Printf("%s is type of S", *start2Fitting.Type)
	var directions = getDirections(start2Fitting)
	log.Printf("Directions: %v", directions)
	var clockwiseCurrent, clockWiseEntry = start2Fitting.getNext(directions[0], s)
	var antiClockwiseCurrent, antiClockwiseEntry = start2Fitting.getNext(directions[1], s)
	var clockWisePosition = clockwiseCurrent.Pos.Offset
	var antiClockwisePosition = antiClockwiseCurrent.Pos.Offset

	collectedPoints[start2Fitting.Pos.Offset] = start2Fitting
	collectedPoints[clockWisePosition] = clockwiseCurrent
	collectedPoints[antiClockwisePosition] = antiClockwiseCurrent

	var distanceCovered = 1
	for clockWisePosition != antiClockwisePosition {
		log.Printf("Step: %d, clock[%d, %d](%s) :%s, anti:[%d, %d](%s): %s", distanceCovered,
			clockwiseCurrent.row(), clockwiseCurrent.column(), clockWiseEntry,
			*clockwiseCurrent.Type,
			antiClockwiseCurrent.row(), antiClockwiseCurrent.column(), antiClockwiseEntry,
			*antiClockwiseCurrent.Type)
		clockwiseCurrent, clockWiseEntry = clockwiseCurrent.getNext(clockWiseEntry, s)
		antiClockwiseCurrent, antiClockwiseEntry = antiClockwiseCurrent.getNext(antiClockwiseEntry, s)
		clockWisePosition = clockwiseCurrent.Pos.Offset
		antiClockwisePosition = antiClockwiseCurrent.Pos.Offset
		collectedPoints[clockWisePosition] = clockwiseCurrent
		collectedPoints[antiClockwisePosition] = antiClockwiseCurrent
		distanceCovered++

		if distanceCovered > MAX_ITERS {
			panic("MAX_ITERS")
		}
	}
	log.Printf("%d is the distance", distanceCovered)
}

func (s *PipeMap) EvaluatePart2() {

	var totalPointsInsideTheLoop = 0
	for _, line := range s.Line {
		var seenPathPointsCount = 0
		for fittingIndex := 0; fittingIndex < len(line.Fitting); {
			var currentFitting = line.Fitting[fittingIndex]
			var isPointInLoop = collectedPoints[currentFitting.Pos.Offset] != nil

			log.Printf("FittingIndex: %d, isPointInLoop: %v", fittingIndex, isPointInLoop)
			if isPointInLoop {
				if currentFitting.opensRight() {

					var directions = getDirections(currentFitting)
					var otherDirection = ""
					var startDirection = ""

					if directions[0] == "right" {
						otherDirection = directions[1]
					} else {
						otherDirection = directions[0]
					}

					startDirection = otherDirection

					var currentNode = &currentFitting

					for currentNode.Pos.Line == currentFitting.Pos.Line {
						fittingIndex += 1

						log.Printf("Navigating loop FittingIndex: %d", fittingIndex)
						currentNode, otherDirection = currentNode.getNext(otherDirection, s)
					}

					var endDirection = ""
					if otherDirection == "top" {
						endDirection = "bottom"
					}
					if otherDirection == "bottom" {
						endDirection = "top"
					}
					if endDirection == startDirection {
						log.Printf("Moving u-turn: %s -> %s", startDirection, endDirection)
						seenPathPointsCount += 2
					} else {
						log.Printf("Moving in-out: %s -> %s", startDirection, endDirection)
						seenPathPointsCount += 1
					}
				} else {
					seenPathPointsCount += 1
					fittingIndex += 1
				}
			} else {
				fittingIndex += 1
				if seenPathPointsCount%2 == 1 {
					log.Printf("Adding fitting with pos [%d, %d], loopCount: %d", currentFitting.Pos.Line, currentFitting.Pos.Column, seenPathPointsCount)
					totalPointsInsideTheLoop++
				}
			}

		}
	}

	log.Printf("%d is the total points inside the loop", totalPointsInsideTheLoop)
}

//TODO: Write a langrange extrapolator value on D3
