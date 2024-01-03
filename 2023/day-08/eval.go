package main

import (
	"github.com/alecthomas/participle/v2"
	"log"
	"regexp"
)

var MAX_ITERS = 1000000

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
	var startNode = nodeMap["AAA"]
	walkNodesPart1(nodeMap, s.DirectionCommand, startNode, func(name string) bool { return name != "ZZZ" })
}

func (s *CamelMap) EvaluatePart2() {

	var startNodes = make([]Node, 0)
	for k, v := range nodeMap {
		found, _ := regexp.MatchString(".*A$", k)
		if found {
			startNodes = append(startNodes, *v)
		}
	}

	steps := make([]int, 0)
	for _, node := range startNodes {
		steps = append(steps, walkNodesPart1(nodeMap, s.DirectionCommand, &node, func(name string) bool {
			found, _ := regexp.MatchString(".*Z$", name)
			return !found
		}))
	}

	log.Printf("%+v", steps)

	var currentLcm = lcm(steps[0], steps[1])

	for i := 2; i < len(steps); i++ {
		currentLcm = lcm(currentLcm, steps[i])
	}

	log.Printf("%d", currentLcm)
}

func lcm(a int, b int) int {
	return a * b / (gcd(a, b))
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func walkNodesPart1(nodeMap map[string]*Node, direction string, startNode *Node, predicate func(name string) bool) int {
	var currentNode = startNode

	var iterationCount = 0
	for predicate(currentNode.Name) {
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
	return iterationCount
}
