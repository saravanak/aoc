package main

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.SimpleRule{

	{"Ident", `[A-Z0-9]+`},
	{"Punct", `[[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	{"EOL", `[\n\r]+`},
	{"whitespace", `[ \t]+`},
})

type CamelMap struct {
	Pos              lexer.Position
	DirectionCommand string     ` @Ident EOL`
	Commands         []*Command `@@*`
}

type Command struct {
	Pos           lexer.Position
	NodesMappings *NodesMapping `@@`
}

type NodesMapping struct {
	Pos    lexer.Position
	Source *Node `@@`
}

type Node struct {
	Name      string  `@Ident`
	Equals    string  `"="`
	Part1     string  `"("`
	Left      *string `@Ident`
	Comma     string  `","`
	Right     *string `@Ident`
	Part2     string  `")"`
	Part3     string  `EOL`
	RightNode *Node
	LeftNode  *Node
}
