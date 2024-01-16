package main

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000

func Parse(contents string) *ProgramText {
	var basicParser = participle.MustBuild[ProgramText](
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

func (s *ProgramText) Evaluate() {
	var sum = 0
	for _, command := range *s.Commands {
		var currentHash = calcWeights(command.Word)
		log.Printf("'%s' :%d", command.Word, currentHash)
		sum += currentHash
	}
	log.Printf("%d", sum)
}

var weights = make(map[string]int)

func calcWeights(input string) int {
	var result = 0

	// log.Printf("%v", weights)
	var foundCachedValue = false
	var currentSearchIndex = len(input)

	for !foundCachedValue && currentSearchIndex >= 0 {
		var searchString = input[0:currentSearchIndex]
		result = weights[searchString]
		foundCachedValue = result > 0
		// log.Printf("%v %d %s", foundCachedValue, result, searchString)
		if !foundCachedValue {
			currentSearchIndex--
		}
	}

	var startIndex = 0
	if foundCachedValue {
		startIndex = currentSearchIndex
	}
	for i := startIndex; i < len(input); i++ {
		var asciiValue = int(input[i])
		result += asciiValue
		result *= 17
		result %= 256
		weights[input[0:i+1]] = result
	}

	return result
}

type LenseSlot struct {
	label string
	fl    int //focallen
}

type Box struct {
	lenses []LenseSlot
}

func (s *ProgramText) EvaluatePart2() {
	s.Evaluate() // prime the maps

	var boxes = make([]Box, 256)

	for _, box := range boxes {
		box.lenses = make([]LenseSlot, 0)
	}
	for _, command := range *s.Commands {
		var currentHash = calcWeights(command.Word)

		var splits = strings.Split(command.Word, "=")

		if len(splits) == 1 {
			var dashSplit = strings.Split(command.Word, "-")
			var boxNumber = weights[dashSplit[0]]

			var label = dashSplit[0]
			var lenses = boxes[boxNumber].lenses
			log.Printf("Remove lense at box: %d label %s", boxNumber, label)

			var lenseWithSameLabel = slices.IndexFunc(lenses, (func(lense LenseSlot) bool {
				return lense.label == label
			}))
			if lenseWithSameLabel >= 0 {
				boxes[boxNumber].lenses = slices.Delete(lenses, lenseWithSameLabel, lenseWithSameLabel+1)
			}
		} else {
			var fl, _ = strconv.Atoi(splits[1])
			var boxNumber = weights[splits[0]]
			var label = splits[0]
			log.Printf("Adding lense to box %d with label %s and fl %d", boxNumber, label, fl)
			var lenses = boxes[boxNumber].lenses

			var lenseWithSameLabel = slices.IndexFunc(lenses, (func(lense LenseSlot) bool {
				return lense.label == label
			}))
			if lenseWithSameLabel >= 0 {
				lenses[lenseWithSameLabel].fl = fl
			} else {
				boxes[boxNumber].lenses = append(lenses, LenseSlot{label, fl})
			}
		}
		log.Printf("'%s' :%d", command.Word, currentHash)
	}
	var sum = 0
	for boxIndex, box := range boxes {
		for lenseIndex, lense := range box.lenses {
			sum += (boxIndex + 1) * (lenseIndex + 1) * lense.fl
		}
	}

	log.Printf("%d", sum)
}
