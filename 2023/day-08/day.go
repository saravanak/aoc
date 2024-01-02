// nolint: golint
package main

/**
*
*  Part 1: 248569531
*  Part 2: 250382098
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"os"
)

func part01(ast *CamelMap) {
	ast.Evaluate()
}

func part02(ast *CamelMap) {
	ast.EvaluatePart2()
}

func main() {

	// fileName := "./data/08/example.txt"
	// fileName := "./data/08/example2.txt"
	fileName := "./data/08/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day08/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day08/part 01!!!")
		part01(ast)
	}

}
