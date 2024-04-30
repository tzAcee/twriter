package compiler

type Parser struct {
	tokenIterator *TokenIterator
}

func NewParser(tIter *TokenIterator) *Parser {
	return &Parser{tIter}
}

func (p *Parser) Parse() (AST, error) {
	return AST{}, nil
}
