// nolint: golint
package main

/**
*
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"os"
)

func part01(ast *RacingRecords) {
	ast.Evaluate()
}

func part02(ast *RacingRecords) {
	ast.EvaluatePart2()
}

func main() {

	fileName := "./data/06/example.txt"
	// fileName := "./data/06/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day06/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day06/part 01!!!")
		part01(ast)
	}

}
