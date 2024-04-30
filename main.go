package main

import (
	"flag"
	"fmt"
	"twriter/compiler"
)

func main() {
	parseExpression := flag.String("exp", "", "expression to parse")
	flag.Parse()
	if len(*parseExpression) == 0 {
		panic("parse expression cannot be empty")
	}
	fmt.Printf("Lexing: '%v'\n", *parseExpression)
	lexemes := compiler.Lex(*parseExpression)

	tokenIterator := compiler.NewTokenIterator(&lexemes)
	parser := compiler.NewParser(tokenIterator)

	ast, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
}
