package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 1000000

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

type Beam struct {
	direction string // rlud
	location  *Point
}
type Point struct {
	x int
	y int
}

var beamMap = make(map[string]int, 0)
var beamTips = make([]Beam, 1)
var visitedLocations = make(map[string]int, 0)

func resetState() {
	//TODO this is a hack: we don't necessarily need to reset the state here . But somehow this is needed for the solution to work
	beamMap = make(map[string]int, 0)
	beamTips = make([]Beam, 1)
	visitedLocations = make(map[string]int, 0)
}

func (s *PipeMap) Evaluate() {
	log.Printf("%v", s.Line)
	var visitedLocations = calcEnergizedCells(s, 0, -1, "r")
	log.Printf("%d", visitedLocations)

}

func (s *PipeMap) EvaluatePart2() {
	var maxEnergizedCellCounts = 0
	for i := 0; i < len(s.Line[0].Places); i++ {
		var cellCounts = calcEnergizedCells(s, -1, i, "d")
		log.Printf("Mving down from top row at pos %d ; count %d", i, cellCounts)
		maxEnergizedCellCounts = max(maxEnergizedCellCounts, cellCounts)
		resetState()
		cellCounts = calcEnergizedCells(s, len(s.Line[0].Places), i, "u")
		log.Printf("Mving up from top row at pos %d ; count %d", i, cellCounts)
		maxEnergizedCellCounts = max(maxEnergizedCellCounts, cellCounts)
		resetState()
	}
	for i := 0; i < len(s.Line); i++ {
		var cellCounts = calcEnergizedCells(s, i, -1, "r")
		log.Printf("Mving right from leftmost col at line %d ; count %d", i, cellCounts)
		resetState()
		maxEnergizedCellCounts = max(maxEnergizedCellCounts, cellCounts)
		cellCounts = calcEnergizedCells(s, i, len(s.Line), "l")
		log.Printf("Mving lett from rightmost col at line %d ; count %d", i, cellCounts)
		resetState()
		maxEnergizedCellCounts = max(maxEnergizedCellCounts, cellCounts)
	}
	log.Printf("%d", maxEnergizedCellCounts)
}

func calcEnergizedCells(s *PipeMap, line int, col int, direction string) int {
	var location = &Point{x: line, y: col}
	var firstBeam = Beam{direction, location}
	beamTips[0] = firstBeam

	var currentIteration = 0
	for len(beamTips) > 0 {
		currentIteration += 1
		var tipBeam = beamTips[len(beamTips)-1]
		// log.Printf("Progressing beam %s", tipBeam.keyFor())
		move(&tipBeam, s)

		beamTips = slices.DeleteFunc(beamTips, func(currentBeam Beam) bool {
			return tipBeam == currentBeam
		})
		visitedLocations[tipBeam.keyForPoint()] = 1
		if currentIteration > MAX_ITERS {
			panic("MAX ITERATIONS")
		}
	}
	return len(visitedLocations) - 1
}

func (b *Beam) keyFor() string {
	return fmt.Sprintf("[%d,%d]->%s", b.location.x, b.location.y, b.direction)
}
func (b *Beam) keyForPoint() string {
	return fmt.Sprintf("[%d,%d]", b.location.x, b.location.y)
}

func move(beam *Beam, s *PipeMap) {
	// log.Printf("At beam %s", beam.keyFor())
	if _, ok := beamMap[beam.keyFor()]; ok {
		//We've visited this one already
		// log.Printf("visited already")
		return
	}

	beamMap[beam.keyFor()] = 1
	// log.Printf("TipsCount: %d. current at %d %d moving %s", len(beamTips), beam.location.x, beam.location.y, beam.direction)

	visitedLocations[beam.keyForPoint()] = 1
	// log.Printf("Incoming direction is %s", beam.direction)

	var mirrorAtPoint Place
	nextPoint := *beam.location

	var nextLine = beam.location.x + rowDelta(beam.direction)
	var nextCol = beam.location.y + columnDelta(beam.direction)
	if nextCol < len(s.Line[0].Places) && nextCol >= 0 &&
		nextLine < len(s.Line) && nextLine >= 0 {
		nextPoint = Point{x: nextLine, y: nextCol}
		mirrorAtPoint = s.Line[nextPoint.x].Places[nextPoint.y]
	} else {
		// log.Printf("Point %d %d outside grid. Stopping beam", nextLine, nextCol)
		return
	}

	if mirrorAtPoint.Type == "." {
		// log.Printf("Seeing . Moving in same direction %v", nextPoint)
		beam.location = &nextPoint
		move(beam, s)
		return
	}

	if canSlide(beam.direction, mirrorAtPoint.Type) {
		// log.Printf("Sliding through..Moving in same direction %v", nextPoint)
		beam.location = &nextPoint
		move(beam, s)
		return
	}
	var reflectedDir, isReflected = reflectedDirection(beam.direction, mirrorAtPoint.Type)

	if isReflected {
		// log.Printf("Reflected. 90 deg rotation..%v  %s", nextPoint, reflectedDir)
		beam.location = &nextPoint
		beam.direction = reflectedDir
		move(beam, s)
		return
	}
	var splitDirs, isSplit = splitDirection(beam.direction, mirrorAtPoint.Type)

	if isSplit {
		// log.Printf("split directions into perperndicular .%v %v", nextPoint, splitDirs)
		beamTips = slices.DeleteFunc(beamTips, func(currentBeam Beam) bool {
			return *beam == currentBeam
		})
		for _, dir := range splitDirs {
			var newBeam = Beam{location: &nextPoint, direction: dir}
			beamTips = append(beamTips, newBeam)
		}
		return
	}

	panic("unmet condition!!!")

}

var columnDelta = func(direction string) int {
	if direction == "r" {
		return +1
	}
	if direction == "l" {
		return -1
	}
	return 0
}
var rowDelta = func(direction string) int {
	if direction == "d" {
		return +1
	}
	if direction == "u" {
		return -1
	}
	return 0
}

var canSlide = func(direction string, mirrorType string) bool {
	if slices.Contains([]string{"r", "l"}, direction) && mirrorType == "-" {
		return true
	}
	if slices.Contains([]string{"u", "d"}, direction) && mirrorType == "|" {
		return true
	}
	return false
}

var reflectedDirection = func(direction string, mirrorType string) (reflectedDir string, isRefelected bool) {

	if mirrorType == "\\" {
		if direction == "r" {
			return "d", true
		}
		if direction == "l" {
			return "u", true
		}
		if direction == "u" {
			return "l", true
		}

		if direction == "d" {
			return "r", true
		}
	}
	if mirrorType == "/" {

		if direction == "r" {
			return "u", true
		}
		if direction == "l" {
			return "d", true
		}
		if direction == "u" {
			return "r", true
		}

		if direction == "d" {
			return "l", true
		}
	}
	return direction, false
}
var splitDirection = func(direction string, mirrorType string) (reflectedDirs []string, isRefelected bool) {

	if mirrorType == "|" {
		if slices.Contains([]string{"r", "l"}, direction) {
			return []string{"d", "u"}, true
		}
	}
	if mirrorType == "-" {
		if slices.Contains([]string{"u", "d"}, direction) {
			return []string{"l", "r"}, true
		}
	}
	return make([]string, 0), false
}
