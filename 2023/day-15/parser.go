package main

import (
	"github.com/alecthomas/participle/v2/lexer"
	// "log"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[^,\n]+`},
	{"EOL", `[\n\r]+`},
	{"Punct", `[,]`},
	{"whitespace", `[ \t]+`},
})

type ProgramText struct {
	Pos      lexer.Position
	Commands *[]Command `@@+`
}

type Command struct {
	Pos   lexer.Position
	Word  string `@Ident`
	Comma string `("," | EOL EOF)`
}
