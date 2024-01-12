package main

import (
	b "aoc/utils"

	"github.com/alecthomas/participle/v2/lexer"
	"log"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[.#]+`},
	{"BlockSeperator", `[\n\r]{2}`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type SpringField struct {
	Pos     lexer.Position
	Pattern []Pattern `@@+`
}

type Pattern struct {
	Pos          lexer.Position
	SpringStatus []PLines `@@+`
	LastLine     LastLine `@@?`
	PatternLines []string
	RuneMatrix   [][]rune
}
type PLines struct {
	SpringSequence string `@Ident`
	LineFeed       string ` EOL (?! @EOL)`
}
type LastLine struct {
	SpringSequence string `@Ident `
	LineFeed       string `@BlockSeperator`
}

func (p *Pattern) join() {
	log.Printf("Joining items")
	p.PatternLines = make([]string, len(p.SpringStatus))
	copy(p.PatternLines, b.Map(p.SpringStatus, (func(pl PLines) string { return pl.SpringSequence })))
	if len(p.LastLine.LineFeed) != 0 {
		p.PatternLines = append(p.PatternLines, p.LastLine.SpringSequence)
	}

	p.RuneMatrix = make([][]rune, 0)
	for _, pattern := range p.PatternLines {
		p.RuneMatrix = append(p.RuneMatrix, []rune(pattern))
	}

}
