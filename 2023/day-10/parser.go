package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[-|LJF7S.]`},
	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type PipeMap struct {
	Pos  lexer.Position
	Line []Line `(@@ EOL)*`
}

type Line struct {
	Pos     lexer.Position
	Fitting []Fitting `@@+`
}

type Fitting struct {
	Pos  lexer.Position
	Type *string `@Ident`
	FittingBehaviour
}
