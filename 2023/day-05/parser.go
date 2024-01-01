package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"MappingHeaderSection", `[a-z]+-to-[a-z]+`},
	{"Ident", `[a-zA-Z_-]+`},
	{"Number", `[-+]?(\d*\.)?\d+`},
	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type SeedingPlan struct {
	Pos      lexer.Position
	Commands []*Command `@@*`
}

type Command struct {
	Pos               lexer.Position
	GivenSeedsCommand *Seeds          ` (  @@ EOL`
	MappingSection    *MappingSection `| @@  ) `
}

type Seeds struct {
	Pos            lexer.Position
	SeedsCommand   string            `"seeds"`
	SeedsSeperator string            `":"`
	SeedsList      []*SeedExpression `@@* `
}

type SeedExpression struct {
	Pos     lexer.Position
	SeedIds *int `@Number`
}

type MappingSection struct {
	Pos              lexer.Position
	MappingHeader    string           `@MappingHeaderSection `
	MapCommandSuffix string           `"map"`
	Seperator        string           `":" EOL`
	MappingLines     []*MappingRecord `@@*`
}

type MappingRecord struct {
	Pos         lexer.Position
	DestStart   *int   `@Number`
	SourceStart *int   `@Number`
	RangeLength *int   `@Number`
	EOL         string `EOL`
}
