// nolint: golint
package main

/**
*
 */
import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"os"
)

func part01(ast *SeedingPlan, context *Context) {
	ast.Evaluate(context)
}

func part02(ast *SeedingPlan, context *Context) {
	ast.EvaluatePart2(context)
}

func main() {

	// fileName := "./data/05/example.txt"
	fileName := "./data/05/full.txt"
	fileContents, _ := script.File(fileName).String()

	ast := Parse(fileContents)

	if ast == nil {
		return
	}
	log.Println(*ast.Commands[0].GivenSeedsCommand.SeedsList[0].SeedIds)
	log.Println(*&ast.Commands[1].MappingSection.MappingHeader)

	var context = Context{seeds: make([]int, 0), mappingRecords: make([]MappingRecord, 0)}

	if os.Args[1] == "part2" {
		fmt.Println("running day05/part 02!!!")
		part02(ast, &context)
	} else {
		fmt.Println("running day05/part 01!!!")
		part01(ast, &context)
	}

}
