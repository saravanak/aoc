package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[A-Z2-9]{5}`},
	{"Number", `[-+]?(\d*\.)?\d+`},
	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type CardsList struct {
	Pos   lexer.Position
	Cards []*Card `@@*`
}

type Card struct {
	Pos          lexer.Position
	CardName     *string `@Ident`
	Bid          *int    `@Number`
	EOL          string  `EOL`
	catetgory    string
	frequencyMap map[rune]int
	sortedName   string
}
