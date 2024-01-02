package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

// var basicLexer = lexer.MustSimple([]lexer.SimpleRule{
//
// 	{"MappingHeaderSection", `[a-z]+-to-[a-z]+`},
// 	{"Ident", `[a-zA-Z_-]+`},
// 	{"Number", `[-+]?(\d*\.)?\d+`},
// 	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
// 	{"EOL", `[\n\r]+`},
// 	{"whitespace", `[ \t]+`},
// })

type RacingRecords struct {
	Pos      lexer.Position
	Commands []*Command `@@*`
}

type Command struct {
	Pos             lexer.Position
	TimeCommand     *Times           ` (  @@`
	DistanceCommand *DistanceCovered `| @@  ) `
}

type Times struct {
	Pos            lexer.Position
	TimesPrefix    string `"Time"`
	TimesDelimiter string `":"`
	RaceTime       *[]int `@Int*`
}

type DistanceCovered struct {
	Pos            lexer.Position
	TimesPrefix    string `"Distance"`
	TimesDelimiter string `":"`
	Distances      *[]int `@Int*`
}
