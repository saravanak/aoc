package main

import (
	// "fmt"
	"fmt"
	"log"

	"github.com/alecthomas/participle/v2"
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

func (s *PipeMap) Evaluate(growthRate int) {

	s.columnThickeness = make(map[int]int)
	s.lineThickness = make(map[int]int)
	log.Printf("%+v", s.Line[0].Places[0].Type)
	columnsCount := len(s.Line[0].Places)
	var columnsHasGalaxies = make([]bool, columnsCount+1)
	var linesWithoutGalaxies = make([]int, 0)

	for lineIndex, line := range s.Line {
		var lineHasGalaxies = false
		for columnIndex, place := range line.Places {
			var isPlaceAGalaxy = place.Type == "#"
			if isPlaceAGalaxy {
				s.GalaxyLocations = append(s.GalaxyLocations, &line.Places[columnIndex])
			}
			if !columnsHasGalaxies[place.Pos.Column] {
				columnsHasGalaxies[place.Pos.Column] = isPlaceAGalaxy
			}
			if !lineHasGalaxies {
				lineHasGalaxies = isPlaceAGalaxy
			}
		}

		if !lineHasGalaxies {
			log.Printf("Line %d HasGalaxies? %v", lineIndex, lineHasGalaxies)
			linesWithoutGalaxies = append(linesWithoutGalaxies, lineIndex)
			s.lineThickness[line.Pos.Line] = growthRate
		}

	}
	for colIndex, colHasGalaxie := range columnsHasGalaxies {
		if !colHasGalaxie {
			log.Printf("Column %d does not have galaxy", colIndex)

			s.columnThickeness[colIndex] = growthRate
		}
	}

	for _, place := range s.GalaxyLocations {
		log.Printf("%d,%d", place.Pos.Line, place.Pos.Column)
	}

	var distances = make(map[*Place]map[*Place]int)
	// for i := 0; i < 1; i++ {
	for i := 0; i < len(s.GalaxyLocations); i++ {
		for j := 0; j < len(s.GalaxyLocations); j++ {
			if i == j {
				continue
			}

			var lhsGalazy = s.GalaxyLocations[i]
			var rhsGalazy = s.GalaxyLocations[j]

			if distances[lhsGalazy] == nil {
				distances[lhsGalazy] = make(map[*Place]int)
			}

			if distances[rhsGalazy] != nil && distances[rhsGalazy][lhsGalazy] > 0 {
				log.Printf("We have got a distance already from [%d,%d] to [%d,%d]", rhsGalazy.Pos.Line, rhsGalazy.Pos.Column, lhsGalazy.Pos.Line, lhsGalazy.Pos.Column)
				continue
			}

			if distances[lhsGalazy][rhsGalazy] == 0 {
				distances[lhsGalazy][rhsGalazy] = 990909090909
			}

			var lineDistance = 0
			var colDistance = 0
			if lhsGalazy.Pos.Line != rhsGalazy.Pos.Line {
				var from = min(lhsGalazy.Pos.Line, rhsGalazy.Pos.Line)
				var to = max(lhsGalazy.Pos.Line, rhsGalazy.Pos.Line)

				log.Printf("line from and to [%d %d]", from, to)
				for k := from + 1; k <= to; k++ {
					log.Printf("line thickness at [%d] is %d", k, s.lineThickness[k])
					if s.lineThickness[k] > 0 {
						lineDistance += s.lineThickness[k]
					} else {
						lineDistance += 1
					}
				}
			}
			if lhsGalazy.Pos.Column != rhsGalazy.Pos.Column {
				var from = min(lhsGalazy.Pos.Column, rhsGalazy.Pos.Column)
				var to = max(lhsGalazy.Pos.Column, rhsGalazy.Pos.Column)

				log.Printf("Col from and to [%d %d]", from, to)
				for k := from + 1; k <= to; k++ {
					log.Printf("Col thickness at [%d] is %d", k, s.columnThickeness[k])
					if s.columnThickeness[k] > 0 {
						colDistance += s.columnThickeness[k]
					} else {
						colDistance += 1
					}
				}
			}

			distances[lhsGalazy][rhsGalazy] = min(
				distances[lhsGalazy][rhsGalazy],
				lineDistance+colDistance,
			)
			log.Printf("Distance Betweeen [%d,%d] and [%d,%d] is %d+%d", lhsGalazy.Pos.Line, lhsGalazy.Pos.Column, rhsGalazy.Pos.Line, rhsGalazy.Pos.Column, lineDistance, colDistance)
		}
	}

	var sumOfShortestDistances = 0
	for k, v := range distances {
		for k1, v1 := range v {
			sumOfShortestDistances += v1
			log.Printf("shortest Distance Betweeen [%d,%d] and [%d,%d] is %d", k.Pos.Line, k.Pos.Column, k1.Pos.Line, k1.Pos.Column, v1)
		}
	}

	fmt.Printf("%d is the sum of all shortest distances", sumOfShortestDistances)

}

func (s *PipeMap) EvaluatePart2() {

}

//TODO: Write a langrange extrapolator value on D3
