package main

import (
	b "aoc/utils"
	"log"

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
	var sum = 0
	for _, pattern := range s.Pattern {
		pattern.join()
		var currrentMatrixToWorkOn = pattern.RuneMatrix
		var reflectedColumnIndex = findReflectionOnColumn(&currrentMatrixToWorkOn)
		log.Printf("%d ", reflectedColumnIndex)

		if reflectedColumnIndex == -1 {
			var transposed = b.Transpose(pattern.RuneMatrix)

			for _, line := range transposed {
				log.Printf("'%s'", string(line))
			}
			var reflectedColumnIndex = findReflectionOnColumn(&transposed)
			log.Printf("%d ", reflectedColumnIndex)
			sum += 100 * reflectedColumnIndex
		} else {
			sum += reflectedColumnIndex
		}
	}

	log.Printf("%d", sum)

}

func (s *SpringField) EvaluatePart2() {

	var sum = 0
	for _, pattern := range s.Pattern {
		pattern.join()
		var currrentMatrixToWorkOn = pattern.RuneMatrix
		var reflectedColumnIndex = findReflectionOnColumnWithSmudge(&currrentMatrixToWorkOn)
		log.Printf("%d ", reflectedColumnIndex)

		if reflectedColumnIndex == -1 {
			var transposed = b.Transpose(pattern.RuneMatrix)

			log.Printf("Trying transposed")
			for _, line := range transposed {
				log.Printf("'%s'", string(line))
			}
			var reflectedColumnIndex = findReflectionOnColumnWithSmudge(&transposed)
			log.Printf("%d ", reflectedColumnIndex)
			if reflectedColumnIndex == -1 {
				panic("Invalid reflection identified")
			}
			sum += 100 * reflectedColumnIndex
		} else {
			sum += reflectedColumnIndex
		}
	}

	log.Printf("%d", sum)
}

func findReflectionOnColumn(inputMatrix *[][]rune) int {

	var leftLineReflections = make([][]rune, len(*inputMatrix))
	var reflectedColumnIndex = -1

	for columnIndex := 0; columnIndex < len((*inputMatrix)[0]); columnIndex++ {
		var isAllColumnsReflected = true
		for lineIndex, currentLine := range *inputMatrix {
			log.Printf("Processing line %s %d %d", string(currentLine), lineIndex, columnIndex)
			var currentLineReflection = leftLineReflections[lineIndex]
			if currentLineReflection == nil {
				currentLineReflection = make([]rune, 0)
			}
			if columnIndex > 0 {
				//Try reflection check
				var smaller = currentLineReflection
				var larger []rune
				var rightReflection = []rune(currentLine[columnIndex:])
				if len(currentLineReflection) > len(rightReflection) {
					smaller = rightReflection
					larger = currentLineReflection
				} else {
					larger = rightReflection
				}

				var allEqual = true
				for smallIndex, smallRune := range smaller {
					if larger[smallIndex] != smallRune {
						allEqual = false
						break
					}
				}

				if isAllColumnsReflected {

					if allEqual {
						// log.Printf("We've got  a reflection on %d %d %s|%s", lineIndex, columnIndex, string(currentLineReflection), rightReflection)
					} else {
						isAllColumnsReflected = false
					}
				}

			}
			leftLineReflections[lineIndex] = append([]rune{rune(currentLine[columnIndex])}, currentLineReflection...)
		}
		if isAllColumnsReflected && columnIndex > 0 {
			log.Printf("%d col is reflected", columnIndex)
			reflectedColumnIndex = columnIndex
			break
		}
	}

	return reflectedColumnIndex
}
func findReflectionOnColumnWithSmudge(inputMatrix *[][]rune) int {

	var leftLineReflections = make([][]rune, len(*inputMatrix))
	var reflectedColumnIndex = -1

	for columnIndex := 0; columnIndex < len((*inputMatrix)[0]); columnIndex++ {
		var isAllColumnsReflected = true
		var seenSmudgedMirror = false
		for lineIndex, currentLine := range *inputMatrix {
			log.Printf("Processing line %s %d %d", string(currentLine), lineIndex, columnIndex)
			var currentLineReflection = leftLineReflections[lineIndex]
			if currentLineReflection == nil {
				currentLineReflection = make([]rune, 0)
			}
			if columnIndex > 0 {
				//Try reflection check
				var smaller = currentLineReflection
				var larger []rune
				var rightReflection = []rune(currentLine[columnIndex:])
				if len(currentLineReflection) > len(rightReflection) {
					smaller = rightReflection
					larger = currentLineReflection
				} else {
					larger = rightReflection
				}

				var allEqual = true
				for smallIndex, smallRune := range smaller {
					if larger[smallIndex] != smallRune {
						if !seenSmudgedMirror {
							seenSmudgedMirror = true
							log.Printf("Seeing smudged mirror at %d", smallIndex)
							continue //Allow one smudge for entire reflection
						}
						allEqual = false
						break
					}
				}

				if isAllColumnsReflected {
					if allEqual {
						// log.Printf("We've got  a reflection on %d %d %s|%s", lineIndex, columnIndex, string(currentLineReflection), rightReflection)
					} else {
						isAllColumnsReflected = false
					}
				}

			}
			leftLineReflections[lineIndex] = append([]rune{rune(currentLine[columnIndex])}, currentLineReflection...)
		}
		log.Printf("colIndex: %d isAllColumnsReflected %v ", columnIndex, isAllColumnsReflected)
		if isAllColumnsReflected && seenSmudgedMirror && columnIndex > 0 {
			log.Printf("%d col is reflected", columnIndex)
			reflectedColumnIndex = columnIndex
			break
		}
	}

	return reflectedColumnIndex
}
