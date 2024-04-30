package compiler

type ASTNode interface {
}

type AST struct {
	head *ASTNode
}

type TrcNode struct {
	bufferNode *BufferNode
}

type BufferNode struct {
	operations []BufferOperationNode
}

type BufferOperationNode struct {
	bufferElement BufferRuneNode
	operations    []BufferOperation
}

type BufferRuneNode struct {
	element rune
}
