package main

/**
* Introduces generics and utils for Map and Filter on a slice.
* Defines methods on Structs and works with pointers
*
* Part 1: 520135
* Part 2:
 */
import (
	"fmt"
	"log"
	"strconv"

	"os"
	"regexp"

	"github.com/bitfield/script"

	b "aoc/utils"
)

type location struct {
	x int
	y int
}

type program struct {
	value  int
	digits int
	at     location
}

type symbol struct {
	at    location
	shape string
}

type schematic struct {
	programs []program
	symbols  []symbol
	//Linewise index of symbols for easy lookup
	symbols_lookup_table  map[int][]*symbol
	programs_lookup_table map[int][]*program
}

func (s *schematic) append_symbol(x int, y int, shape string) {
	new_symbol := symbol{shape: shape, at: location{x, y}}
	s.symbols = append(s.symbols, new_symbol)
	if s.symbols_lookup_table[y] == nil {
		s.symbols_lookup_table[y] = make([]*symbol, 1)
	}
	s.symbols_lookup_table[y] = append(s.symbols_lookup_table[y], &new_symbol)
}

func (s *schematic) appendProgram(x int, y int, programValue int, digits int) {
	new_program := program{value: programValue, digits: digits, at: location{x, y}}
	s.programs = append(s.programs, new_program)
	if s.programs_lookup_table[y] == nil {
		s.programs_lookup_table[y] = make([]*program, 1)
	}
	s.programs_lookup_table[y] = append(s.programs_lookup_table[y], &new_program)
}

func (thisSymbol *symbol) FindMatchingPrograms() int {
	neighbourLines := []int{(thisSymbol.at.y) - 1, thisSymbol.at.y, (thisSymbol.at.y) + 1}
	neighbourLines = b.Filter(neighbourLines, (func(value int) bool {
		return value >= 0 && value < line_number
	}))

	programsInLine := make([]*program, 1)

	for _, neighbourLine := range neighbourLines {
		programsInLine = append(programsInLine, given_schematic.programs_lookup_table[neighbourLine]...)
	}

	programsInLine = b.Filter(programsInLine, (func(thisProgram *program) bool { return thisProgram != nil }))

	programValues := b.Map(programsInLine, (func(thisProgram *program) program { return *thisProgram }))

	log.Printf("Nearby programs in line for gear %d", thisSymbol.at)
	log.Println(programValues)

	numberOfAdjacentPrograms := 0
	gearRatio := 1
	for _, nearbyProgram := range programValues {
		if thisSymbol.at.x >= nearbyProgram.at.x-1 && thisSymbol.at.x <= nearbyProgram.at.x+nearbyProgram.digits {
			numberOfAdjacentPrograms++
			gearRatio *= nearbyProgram.value

		}
	}
	if numberOfAdjacentPrograms == 2 {
		return gearRatio
	} else {
		return 0
	}

}

func (p *program) FindMatchingSymbols() int {
	neighbourLines := []int{(p.at.y) - 1, p.at.y, (p.at.y) + 1}
	neighbourLines = b.Filter(neighbourLines, (func(value int) bool {
		return value >= 0 && value < line_number
	}))

	symbolsInLine := make([]*symbol, 1)

	for _, neighbourLine := range neighbourLines {
		symbolsInLine = append(symbolsInLine, given_schematic.symbols_lookup_table[neighbourLine]...)
	}

	symbolsInLine = b.Filter(symbolsInLine, (func(s *symbol) bool { return s != nil }))

	symbolValues := b.Map(symbolsInLine, (func(s *symbol) symbol { return *s }))

	log.Printf("Nearby symbols in line for program %d", p.value)
	log.Println(symbolValues)

	hasMatchingSymbol := false
	for _, nearbySymbol := range symbolValues {
		if nearbySymbol.at.x >= p.at.x-1 && nearbySymbol.at.x <= p.at.x+p.digits {
			hasMatchingSymbol = true

			if hasMatchingSymbol {
				return p.value
			}
		}
	}
	if !hasMatchingSymbol {
		log.Printf("Program %d does not have a matching symbol", p.value)
		return 0
	} else {
		return p.value
	}

}

var given_schematic = schematic{symbols_lookup_table: make(map[int][]*symbol), programs_lookup_table: make(map[int][]*program)}
var line_number = 0

func parsePartsAndSymbols(line string, line_number int) {
	numberMatcher := regexp.MustCompile("((?P<Digit>\\d+)|(?P<Special>[^\\d|\\.]))")

	matches := numberMatcher.FindAllStringIndex(line, -1)

	fmt.Println(matches)
	fmt.Println(numberMatcher.SubexpNames())

	for _, index := range matches {
		match := line[index[0]:index[1]]

		programValue, err := strconv.Atoi(match)
		if err == nil {
			given_schematic.appendProgram(index[0], line_number, programValue, index[1]-index[0])
		} else {
			given_schematic.append_symbol(index[0], line_number, match)
			// given_schematic.symbols = append(given_schematic.symbols, this_symbol)
		}

	}
}

func lineParser(line string) string {
	parsePartsAndSymbols(line, line_number)
	line_number++
	return line
}

func part01() {

	log.Printf("total lines : %d", line_number)
	sum := 0
	for _, program := range given_schematic.programs {

		sum = sum + program.FindMatchingSymbols()
	}

	// log.Println(given_schematic)
	fmt.Printf("Day 03 Part 02: %d", sum)
}

func part02() {
	sum := 0
	gears := b.Filter(given_schematic.symbols, (func(s symbol) bool { return s.shape == "*" }))
	for _, gear := range gears {
		sum = sum + gear.FindMatchingPrograms()
	}

	// log.Println(given_schematic)
	fmt.Printf("Day 03 Part 01: %d", sum)
}

func main() {
	fileName := "./data/03/full.txt"
	script.File(fileName).FilterLine(lineParser).Wait()
	if os.Args[1] == "part2" {
		fmt.Println("running day03/part 02!!!")
		part02()
	} else {
		fmt.Println("running day03/part 01!!!")
		part01()
	}
	// log.Println(contents)

}
