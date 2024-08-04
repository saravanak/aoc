package main

import (
	"fmt"
	"log"
	"maps"
	"math"
	"slices"

	"github.com/Goldziher/go-utils/sliceutils"
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
	cost int
	x    int
	y    int
}
type PathNode struct {
	node                *Node
	direction           string
	lastThreeDirections []string
	cumulativeCost      int
	seenNodes           map[string]int
	pathTaken           []string
}

type Point struct {
	x int
	y int
}

var lowestCost = int(math.Pow(10, 10))
var lowestCostNode PathNode
var lineCount = 0
var columnCount = 0
var iterations = 0
var rejectedMap = make(map[string]int)

func (s *PipeMap) Evaluate() {
	lineCount = len(s.Line)
	columnCount = len(s.Line[0].Places)
	// lineCount = 7
	// columnCount = 7
	var destinationNode = Point{x: lineCount - 1, y: columnCount - 1}

	log.Printf("%v", destinationNode)
	log.Printf("%v %d", s.Line[0], lowestCost)
	var nodes = make([][]Node, 0)

	for i := 0; i < lineCount; i++ {
		nodes = append(nodes, make([]Node, 0))
		for j := 0; j < columnCount; j++ {
			var n = Node{cost: s.Line[i].Places[j].Type, x: i, y: j}
			nodes[i] = append(nodes[i], n)
		}
	}
	var simpleLadderCost = 0
	for i := 0; i < lineCount; i++ {
		if i == lineCount-1 {
			simpleLadderCost += nodes[i][i].cost
			continue
		}
		simpleLadderCost += (nodes[i+1][i].cost + nodes[i][i+1].cost)
	}

	lowestCost = simpleLadderCost

	log.Printf("simpleLadder  cost %d", lowestCost)

	s.searchLowestHeat(&nodes, PathNode{node: &nodes[0][0], direction: "", seenNodes: make(map[string]int)})

	log.Printf("%v", lowestCostNode)
}

func Destination(nodes *[][]Node) *Node {
	return &(*nodes)[lineCount-1][columnCount-1]
}

func DebugStr(currentNode PathNode) string {
	var nodeAsString = fmt.Sprintf("%d:%d-%s-%d", currentNode.node.x, currentNode.node.y, currentNode.direction, currentNode.cumulativeCost)
	return nodeAsString
}

func (s *PipeMap) searchLowestHeat(nodesList *[][]Node, currentNode PathNode) {

	iterations += 1
	if reject(nodesList, currentNode) {
		return
	}

	if accept(nodesList, currentNode) {
		return
	}

	log.Printf("%d", iterations)
	if iterations > MAX_ITERS {
		panic("MAX_ITERS")
	}
	for _, nextNode := range nextNodes(nodesList, currentNode) {
		s.searchLowestHeat(nodesList, nextNode)
	}

}

func reject(nodesList *[][]Node, currentNode PathNode) bool {
	if currentNode.direction == "" {
		return false // we are not done with the root node
	}

	var nodeAsString = DebugStr(currentNode)
	if currentNode.seenNodes[nodeAsString] == 1 {
		log.Printf("detected a loop %v", currentNode.pathTaken)
		rejectedMap[nodeAsString] = 1
		return true
	} else {
		currentNode.seenNodes[nodeAsString] = 1
	}

	if currentNode.cumulativeCost > lowestCost {
		log.Printf("Rejecting node at %v %d > %d", currentNode.node, currentNode.cumulativeCost, lowestCost)
		rejectedMap[nodeAsString] = 1
		return true
	}

	var linesRemaining = lineCount - currentNode.node.x - 1
	var columnsRemaining = columnCount - currentNode.node.y - 1

	if currentNode.cumulativeCost+(linesRemaining+columnsRemaining)*3 > lowestCost {
		log.Printf("Rejecting node using manning distance hueristic %d+%d+%d > %d", currentNode.cumulativeCost, linesRemaining, columnsRemaining, lowestCost)
		rejectedMap[nodeAsString] = 1
		return true
	}

	return false
}

func accept(nodesList *[][]Node, currentNode PathNode) bool {
	if currentNode.cumulativeCost == 0 {
		return false
	}
	if currentNode.node == Destination(nodesList) {
		if currentNode.cumulativeCost < lowestCost {
			lowestCostNode = currentNode
			lowestCost = currentNode.cumulativeCost
			log.Printf("Setting new lowest scoere ast %d %v", currentNode.cumulativeCost, currentNode.pathTaken)
		}
		return true
	}
	return false
}

func nextNodes(nodesList *[][]Node, currentNode PathNode) []PathNode {
	// log.Printf("Calling next nodes")
	var result = make([]PathNode, 0)
	if currentNode.direction == "" {
		//root condition
		var intitialDirections = []string{"r", "d"}
		for i := 0; i < len(intitialDirections); i++ {
			result = append(result, PathNode{
				node:                currentNode.node,
				direction:           intitialDirections[i],
				seenNodes:           maps.Clone(currentNode.seenNodes),
				pathTaken:           slices.Clone(currentNode.pathTaken),
				lastThreeDirections: []string{"", "", intitialDirections[i]}})
		}
		return result
	}

	var oppositeDirections = map[string]string{
		"r": "l",
		"l": "r",
		"u": "d",
		"d": "u",
	}

	var directionsToRemove = []string{oppositeDirections[currentNode.direction]}
	var isAllSameDirection = sliceutils.Every(currentNode.lastThreeDirections, func(direction string, index int, directions []string) bool {
		return direction == directions[0]
	})

	if isAllSameDirection {
		directionsToRemove = append(directionsToRemove, currentNode.lastThreeDirections[0])
	}

	for key := range oppositeDirections {

		if slices.Contains(directionsToRemove, key) {
			continue
		}

		var nextNode = currentNode.node

		switch key {
		case "r":
			var nextY = currentNode.node.y + 1
			if nextY >= columnCount {
				continue
			}
			nextNode = &(*nodesList)[currentNode.node.x][nextY]
		case "l":
			var nextY = currentNode.node.y - 1
			if nextY < 0 {
				continue
			}
			nextNode = &(*nodesList)[currentNode.node.x][nextY]
		case "d":
			var nextX = currentNode.node.x + 1
			if nextX >= lineCount {
				continue
			}
			nextNode = &(*nodesList)[nextX][currentNode.node.y]
		case "u":
			var nextX = currentNode.node.x - 1
			if nextX < 0 {
				continue
			}
			nextNode = &(*nodesList)[nextX][currentNode.node.y]
		}
		var lastThreeDirections = append(slices.Clone(currentNode.lastThreeDirections[1:]), key)

		var nodeAsString = DebugStr(currentNode)
		var newPathNode = PathNode{
			node:                nextNode,
			direction:           key,
			cumulativeCost:      currentNode.cumulativeCost + nextNode.cost,
			lastThreeDirections: lastThreeDirections,
			pathTaken:           append(currentNode.pathTaken, nodeAsString),
			seenNodes:           maps.Clone(currentNode.seenNodes),
		}
		if rejectedMap[DebugStr(newPathNode)] == 1 {
			continue
		}
		result = append(result, newPathNode)
	}

	// log.Printf("%v", result)
	// for _, resultNode := range result {
	// 	log.Printf("%+v %v [%d,%d]%s %v %v", resultNode.pathTaken, resultNode.cumulativeCost, resultNode.node.x, resultNode.node.y, resultNode.direction, resultNode.lastThreeDirections, currentNode.lastThreeDirections)
	// }
	return result
}

func (s *PipeMap) EvaluatePart2() {

}
