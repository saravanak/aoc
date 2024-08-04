// nolint: golint
package main

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

	fileName := "./data/17/example.txt"
	// fileName := "./data/17/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}

	if os.Args[1] == "part2" {
		fmt.Println("running day17/part 02!!!")
		part02(ast)
	} else {
		fmt.Println("running day17/part 01!!!")
		part01(ast)
	}

}
