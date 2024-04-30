package compiler

import (
	"errors"
	"fmt"
)

type Parser struct {
	tokenIterator *TokenIterator
}

func NewParser(tIter *TokenIterator) *Parser {
	return &Parser{tIter}
}

func (p *Parser) Parse() (AST, error) {
	trcNode, ok, err := p.parseTrc()
	if !ok {
		return AST{}, err
	} else if ok && err == nil {
		return AST{&trcNode}, nil
	}

	bufferNode, err := p.parseBuffer()
	if err != nil {
		return AST{}, err
	}

	return AST{&bufferNode}, nil
}

// can be optional, so the bool flag is needed
func (p *Parser) parseTrc() (ASTNode, bool, error) {
	// check 'trc'
	// recover if first letter is not 't'
	isRune, err := p.isRune('t')
	if err != nil {
		return TrcNode{}, false, err
	} else if !isRune {
		// recover
		tokenMeta, ok := p.tokenIterator.Get()
		if !ok {
			return TrcNode{}, false, errors.New("token iterators upper limit reached")
		}
		return TrcNode{}, true, fmt.Errorf("expected 't', but got '%v' @[ln %d : col %d]", tokenMeta.Rune(), tokenMeta.Ln(), tokenMeta.Col())
	}
	steps, err := p.isRuneSequence("trc")
	if err != nil {
		return TrcNode{}, false, err
	}
	p.tokenIterator.Step(steps)

	bufferASTNode, err := p.parseBuffer()
	if err != nil {
		return TrcNode{}, false, err
	}

	bufferNode, ok := bufferASTNode.(BufferNode)
	if !ok {
		return TrcNode{}, false, errors.New("could not create buffer node")
	}

	return TrcNode{&bufferNode}, true, nil
}

func (p *Parser) parseBuffer() (ASTNode, error) {
	bufferOperations := []BufferOperationNode{}

	for bufferOpASTNode, ok, err := p.parseBufferOperationNode(); ok; bufferOpASTNode, ok, err = p.parseBufferOperationNode() {
		if err != nil {
			return BufferNode{}, err
		}

		bufferOp, okConv := bufferOpASTNode.(BufferOperationNode)
		if !okConv {
			return BufferNode{}, errors.New("could not create BufferOperationNode")
		}
		bufferOperations = append(bufferOperations, bufferOp)
	}

	return BufferNode{bufferOperations, nil}, nil
}

func (p *Parser) parseBufferOperationNode() (ASTNode, bool, error) {
	var bufferNodeOp BufferOperationNode

	// cant start with bufferOperation
	if isBuffer, err := p.isBufferOperation(); err == nil && isBuffer {
		tokenMeta, ok := p.tokenIterator.Get()
		if !ok {
			return bufferNodeOp, false, errors.New("token iterators upper limit reached")
		}
		return bufferNodeOp, false, fmt.Errorf("got buffer operation, but buffer rune was expected @[ln %d : col %d]", tokenMeta.Ln(), tokenMeta.Col())
	} else if err != nil {
		return bufferNodeOp, false, err
	}

	runeASTNode, err := p.parseBufferRuneNode()
	if err != nil {
		return bufferNodeOp, false, err
	}
	runeNode, ok := runeASTNode.(BufferRuneNode)
	if !ok {
		return bufferNodeOp, false, errors.New("could not create BufferRuneNode")
	}

	bufferOpNodes := []BufferOperation{}
	for {
		bufferOp, ok := p.parseBufferOperation()
		if !ok {
			break
		}

		bufferOpNodes = append(bufferOpNodes, bufferOp)
	}

	bufferNodeOp.bufferElement = runeNode
	bufferNodeOp.operations = bufferOpNodes

	return bufferNodeOp, true, nil
}

func (p *Parser) parseBufferOperation() (BufferOperation, bool) {
	isBufferOp, err := p.isBufferOperation()
	if !isBufferOp || err != nil {
		return 0, false
	}

	tokenMeta, _ := p.tokenIterator.Get()
	p.tokenIterator.Step(1)
	return BufferOperation(tokenMeta.token), true
}

func (p *Parser) isBufferOperation() (bool, error) {
	tokenMeta, ok := p.tokenIterator.Get()
	if !ok {
		return false, errors.New("token iterators upper limit reached")
	}

	ops := [4]Token{
		T_ArrowLeft,
		T_ArrowRight,
		T_CurlyClose,
		T_CurlyOpen,
	}

	for _, op := range ops {
		if op == tokenMeta.token {
			return true, nil
		}
	}

	return false, nil
}

func (p *Parser) parseBufferRuneNode() (ASTNode, error) {
	tokenMeta, ok := p.tokenIterator.Get()
	if !ok {
		return false, errors.New("token iterators upper limit reached")
	}
	if tokenMeta.token == T_Letter || tokenMeta.token == T_Number {
		p.tokenIterator.Step(1)
		return BufferRuneNode{tokenMeta.Rune()}, nil
	}

	return BufferNode{}, fmt.Errorf("expected buffer rune, but got '%v' @[ln %d : col %d]", tokenMeta.token, tokenMeta.Ln(), tokenMeta.Col())
}

// return steps or error
func (p *Parser) isRuneSequence(seq string) (uint16, error) {
	for i, run := range seq {
		tokenMeta, ok := p.tokenIterator.Peek(uint16(i))
		if !ok {
			return 0, errors.New("token iterators upper limit reached")
		}

		if tokenMeta.Rune() != run {
			return 0, fmt.Errorf("expected char '%v', but got '%v' @[ln %d : col %d]", run, tokenMeta.Rune(), tokenMeta.Ln(), tokenMeta.Col())
		}
	}

	return uint16(len(seq)), nil
}

func (p *Parser) isRune(r rune) (bool, error) {
	tokenMeta, ok := p.tokenIterator.Get()
	if !ok {
		return false, errors.New("token iterators upper limit reached")
	}

	if tokenMeta.Rune() == r {
		return true, nil
	}

	return false, nil
}
