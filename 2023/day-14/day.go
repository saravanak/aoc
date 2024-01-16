// nolint: golint
package main

/**
*
*  Part 1: 6838
*  Part 2: 451
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"os"
)

func part01(ast *PipeMap) {
	ast.Evaluate()
}

func part02(ast *PipeMap) {
	ast.EvaluatePart2()
}

func main() {

	fileName := "./data/14/example.txt"
	// fileName := "./data/14/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day14/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day14/part 01!!!")
		part01(ast)
	}

}
