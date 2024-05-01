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
	operations []*BufferOperationNode
	buffer     Buffer
}

func (bn *BufferNode) ToString() string {
	for _, op := range bn.operations {
		op.loadElement(&bn.buffer)
	}

	for _, op := range bn.operations {
		op.execute(&bn.buffer)
	}

	return string(bn.buffer)
}

type BufferOperationNode struct {
	bufferElement BufferRuneNode
	bufferIndex   int
	operations    []BufferOperation
}

func (bon *BufferOperationNode) loadElement(buffer *Buffer) {
	*buffer = append(*buffer, bon.bufferElement.element)
	bon.bufferIndex = len(*buffer)
	if bon.bufferIndex != 0 {
		bon.bufferIndex--
	}
}

func (bon *BufferOperationNode) execute(buffer *Buffer) {

	for _, op := range bon.operations {
		switch op {
		case BufferOperation(T_ArrowLeft):
			bon.shiftLeft(buffer)
		case BufferOperation(T_ArrowRight):
			bon.shiftRight(buffer)
		case BufferOperation(T_CurlyOpen):
			bon.switchLeft(buffer)
		case BufferOperation(T_CurlyClose):
			bon.switchRight(buffer)
		default:
			panic(fmt.Sprintf("unsupported buffer operation %d", op))
		}
	}
}

func (bon *BufferOperationNode) shiftLeft(buffer *Buffer) {
	(*buffer)[bon.bufferIndex] = ' '

	if bon.bufferIndex == 0 {
		*buffer = append([]rune{bon.bufferElement.element}, *buffer...)
	} else {
		newIndex := bon.bufferIndex - 1
		(*buffer)[newIndex] = bon.bufferElement.element
		bon.bufferIndex = newIndex
	}
}

func (bon *BufferOperationNode) switchLeft(buffer *Buffer) {
	if bon.bufferIndex == 0 {
		*buffer = append([]rune{bon.bufferElement.element}, *buffer...)
		(*buffer)[1] = ' '
	} else {
		newIndex := bon.bufferIndex - 1
		tmpVal := (*buffer)[newIndex]
		(*buffer)[newIndex] = bon.bufferElement.element
		(*buffer)[bon.bufferIndex] = tmpVal
		bon.bufferIndex = newIndex
	}
}

func (bon *BufferOperationNode) switchRight(buffer *Buffer) {
	newIndex := bon.bufferIndex + 1
	if bon.bufferIndex == len(*buffer)-1 {
		*buffer = append(*buffer, bon.bufferElement.element)
		(*buffer)[len(*buffer)-2] = ' '
	} else {

		tmpVal := (*buffer)[newIndex]
		(*buffer)[newIndex] = bon.bufferElement.element
		(*buffer)[bon.bufferIndex] = tmpVal
	}

	bon.bufferIndex = newIndex
}

func (bon *BufferOperationNode) shiftRight(buffer *Buffer) {
	(*buffer)[bon.bufferIndex] = ' '
	newIndex := bon.bufferIndex + 1

	if bon.bufferIndex == len(*buffer)-1 {
		*buffer = append(*buffer, bon.bufferElement.element)
	} else {
		(*buffer)[newIndex] = bon.bufferElement.element
	}
	bon.bufferIndex = newIndex
}

type BufferRuneNode struct {
	element rune
}
