package main

import (
	// b "aoc/utils"
	// "fmt"
	"log"
	// "slices"
	// "strconv"

	// "strconv"
	// "strings"

	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 100000

func Parse(contents string) *CamelMap {
	var basicParser = participle.MustBuild[CamelMap](
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

func (s *CamelMap) Evaluate() {
	log.Printf("%+v", nodeMap)
	walkNodesPart1(nodeMap, s.DirectionCommand)
}

func (s *CamelMap) EvaluatePart2() {

}

func walkNodesPart1(nodeMap map[string]*Node, direction string) {
	var startNode = nodeMap["AAA"]
	var currentNode = startNode

	var iterationCount = 0
	for currentNode.Name != "ZZZ" {
		if iterationCount > MAX_ITERS {
			log.Printf("MAX ITERATION EXCEEDED")
			break
		}
		var currentDirection = direction[iterationCount%len(direction)]
		switch rune(currentDirection) {
		case 'R':
			log.Printf("Moving R from %+v", currentNode.Name)
			currentNode = nodeMap[*currentNode.Right]
			log.Printf("to %+v", currentNode.Name)
		case 'L':
			log.Printf("Moving L from %+v", currentNode.Name)
			currentNode = nodeMap[*currentNode.Left]
			log.Printf("to %+v\n", currentNode.Name)
		}
		iterationCount += 1
	}
	log.Printf("Journey took %d Steps", iterationCount)
}
