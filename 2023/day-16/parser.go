package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[.|\\/-]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type PipeMap struct {
	Pos  lexer.Position
	Line []Line `(@@ EOL)*`
}

type Line struct {
	Pos    lexer.Position
	Places []Place `@@+`
}

type Place struct {
	Pos  lexer.Position
	Type string `@Ident`
}
