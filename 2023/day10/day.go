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

	// fileName := "./data/10/example.txt"
	// fileName := "./data/10/example2.txt"
	// fileName := "./data/10/example3.txt"
	// fileName := "./data/10/example4.txt"
	// fileName := "./data/10/example5.txt"
	// fileName := "./data/10/example6.txt"
	// fileName := "./data/10/example7.txt"
	// fileName := "./data/10/example8.txt"
	fileName := "./data/10/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day10/part 02!!!")
		part01(ast)
		part02(ast)
	} else {
		fmt.Println("running day10/part 01!!!")
		part01(ast)
	}

}
