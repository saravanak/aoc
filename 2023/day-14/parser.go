package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[.#O]+`},
	{"BlockSeperator", `[\n\r]{2}`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type PipeMap struct {
	Pos        lexer.Position
	Line       []Line `(@@ EOL)*`
	RuneMatrix [][]rune
}

type Line struct {
	Pos    lexer.Position
	Places string `@Ident`
}

func (p *PipeMap) join() {
	p.RuneMatrix = make([][]rune, 0)
	for _, line := range p.Line {
		p.RuneMatrix = append(p.RuneMatrix, []rune(line.Places))
	}
}

func (s *PipeMap) calcLoads() int {
	var sum = 0
	for lineIndex, line := range s.RuneMatrix {
		for _, stone := range line {
			if stone == 'O' {
				sum += len(s.RuneMatrix) - lineIndex
			}
		}
	}
	return sum

}
