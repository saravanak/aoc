// nolint: golint
package main

/**
*
*  Part 1: 1861775706
*  Part 2: 1082
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"os"
)

func part01(ast *OasisReadings) {
	ast.Evaluate()
}

func part02(ast *OasisReadings) {
	ast.EvaluatePart2()
}

func main() {

	// fileName := "./data/09/example.txt"
	fileName := "./data/09/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day09/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day09/part 01!!!")
		part01(ast)
	}

}
