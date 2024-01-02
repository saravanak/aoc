// nolint: golint
package main

/**
*
*  Part 1:
*  Part 2:
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"os"
)

func part01(ast *CardsList) {
	ast.Evaluate()
}

func part02(ast *CardsList) {
	ast.EvaluatePart2()
}

func main() {

	fileName := "./data/07/example.txt"
	// fileName := "./data/07/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day07/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day07/part 01!!!")
		part01(ast)
	}

}
