package compiler

import "fmt"

type Buffer []rune

type ASTNode interface {
}

type AST struct {
	head *ASTNode
}

type TrcNode struct {
	bufferNode *BufferNode
}

func (tn *TrcNode) Trace() {
	buffer := tn.bufferNode.ToString()
	fmt.Printf("buffer sz: %d\n", len(buffer))
	fmt.Printf("[%v]\n", buffer)
}

type BufferNode struct {
	operations []BufferOperationNode
	buffer     Buffer
}

func (bn *BufferNode) ToString() string {
	for _, op := range bn.operations {
		op.execute(&bn.buffer)
	}

	return string(bn.buffer)
}

type BufferOperationNode struct {
	bufferElement BufferRuneNode
	operations    []BufferOperation
}

func (bon *BufferOperationNode) execute(buffer *Buffer) {
	*buffer = append(*buffer, bon.bufferElement.element)

	currentIndex := len(*buffer)
	if currentIndex != 0 {
		currentIndex--
	}

	for _, op := range bon.operations {
		switch op {
		case BufferOperation(T_ArrowLeft):
			bon.shiftLeft(&currentIndex, buffer)
		case BufferOperation(T_ArrowRight):
			bon.shiftRight(&currentIndex, buffer)
		default:
			panic(fmt.Sprintf("unsupported buffer operation %d", op))
		}
	}
}

func (bon *BufferOperationNode) shiftLeft(currentIndex *int, buffer *Buffer) {
	(*buffer)[*currentIndex] = ' '

	if *currentIndex == 0 {
		*buffer = append([]rune{bon.bufferElement.element}, *buffer...)
	} else {
		newIndex := *currentIndex - 1
		(*buffer)[newIndex] = bon.bufferElement.element
		*currentIndex = newIndex
	}

}

func (bon *BufferOperationNode) shiftRight(currentIndex *int, buffer *Buffer) {
	(*buffer)[*currentIndex] = ' '

	if *currentIndex == len(*buffer)-1 {
		*buffer = append(*buffer, bon.bufferElement.element)
	} else {
		newIndex := *currentIndex + 1
		(*buffer)[newIndex] = bon.bufferElement.element
		*currentIndex = newIndex
	}
}

type BufferRuneNode struct {
	element rune
}
