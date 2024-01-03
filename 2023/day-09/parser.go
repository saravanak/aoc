package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[0-9-]+`},
	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type OasisReadings struct {
	Pos      lexer.Position
	Readings []*ReadingLine `(@@ EOL)*`
}

type ReadingLine struct {
	Pos     lexer.Position
	Reading *[]int `@Ident+`
}
