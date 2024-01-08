package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[.#?]`},
	{"Number", `[\d]`},
	{"Punct", `,`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type SpringField struct {
	Pos          lexer.Position
	SpringStatus []SpringStatus `(@@ EOL)*`
}

type SpringStatus struct {
	Pos            lexer.Position
	SpringSequence string     `@Ident*`
	Checksum       []Checksum `@@*`
}

type Checksum struct {
	Pos            lexer.Position
	SingleChecksum int    `@Number`
	Comma          string `(",")?`
}
