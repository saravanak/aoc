package main

import (
	"fmt"
	"log"
	"strings"

	// "maps"
	// "math"
	// "slices"
	//
	// "github.com/Goldziher/go-utils/sliceutils"
	"github.com/alecthomas/participle/v2"
)

var MAX_ITERS = 10000000

// var MAX_ITERS = 5000

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

type Node struct {
	cost              int
	x                 int
	y                 int
	allowedDirections []string
	paths             map[string]int
}

type Point struct {
	x int
	y int
}

func Keys[V any](m map[string]V) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

var totalRows int
var totalCols int
var nodes = make([][]Node, 0)

func (s *PipeMap) Evaluate() {
	lineCount := len(s.Line)
	columnCount := len(s.Line[0].Places)
	// lineCount = 7
	// columnCount = 7
	//
	totalCols = columnCount
	totalRows = lineCount

	fmt.Printf("%d:%d", lineCount, columnCount)

	for i := 0; i < lineCount; i++ {
		nodes = append(nodes, make([]Node, 0))
		for j := 0; j < columnCount; j++ {
			var n = Node{cost: s.Line[i].Places[j].Type, x: i, y: j, allowedDirections: make([]string, 0)}
			nodes[i] = append(nodes[i], n)
		}
	}

	for i := 0; i < lineCount; i++ {
		for j := 0; j < columnCount; j++ {
			node := nodes[i][j]
			if j > 0 {
				node.allowedDirections = append(node.allowedDirections, "t")
			}
			if i > 0 {
				node.allowedDirections = append(node.allowedDirections, "l")
			}
			if i < lineCount-1 {
				node.allowedDirections = append(node.allowedDirections, "r")
			}
			if j < columnCount-1 {
				node.allowedDirections = append(node.allowedDirections, "b")
			}
			nodes[i][j] = node
			fmt.Printf("(%d,%d):%v\n", i, j, node.allowedDirections)
		}
	}

	fmt.Printf("Distances %v\n", nodes[0][0].allowedDirections)

	walk(&nodes[0][0], 0, "")

	// minCost := 10000

	// for dir, node := range nodes[1][0].next {
	// 	fmt.Printf("dir: %v, Cost:  %d\n", dir, node.cost)
	// 	minCost = min(node.cost, minCost)
	// }

	// fmt.Printf("Min distance: %d", minCost)

}

func walk(currentNode *Node, cost int, directionStack string) (string, int) {
	for _, dir := range currentNode.allowedDirections {
		// fmt.Printf("%v", dir)

		for suffixes := range string["rrr", "lll", "bbb", "ttt"] {

			strings.HasSuffix(directionStack, suffix) 
		}


		if dir == "r" {
			key, cost := walk(&nodes[currentNode.x][currentNode.y+1], cost+currentNode.cost, directionStack+"r")
			currentNode.paths[key] = cost
		}
		if dir == "t" {
			nextNode = &(*nodes)[currentNode.x][currentNode.y-1]
			oppositeDirection = "b"
		}
		if dir == "b" {
			nextNode = &(*nodes)[currentNode.x][currentNode.y+1]
			oppositeDirection = "t"
		}
		if dir == "l" {
			nextNode = &(*nodes)[currentNode.x-1][currentNode.y]
			oppositeDirection = "r"
		}

		// fmt.Printf("%s from (%d,%d) to (%d, %d)\n", dir, currentNode.x, currentNode.y, nextNode.x, nextNode.y)
		// if nodeForDirection, ok := currentNode.next[dir]; !ok {
		// 	currentNode.next[dir] = NextNode{node: nextNode, cost: currentCost}
		// 	fmt.Printf("\tCost : %d\n", currentCost)
		// } else if currentCost < nodeForDirection.cost {
		// 	currentNode.next[dir] = NextNode{node: nextNode, cost: currentCost}
		// 	fmt.Printf("\tCost : %d\n", currentCost)
		// } else {
		// 	fmt.Printf("using old values\n")
		// }

		// if _, ok := currentNode.next[dir].node.next[oppositeDirection]; !ok {
		// 	walk(nextNode, nodes, currentCost)
		// }

	}

}

func (s *PipeMap) EvaluatePart2() {

}
