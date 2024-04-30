package compiler

func Traverse(ast *AST) *string {
	switch node := (*ast.head).(type) {
	case BufferNode:
		buffer := node.ToString()
		return &buffer
	case TrcNode:
		node.Trace()
	default:
		panic("unexpected AST node")
	}

	return nil
}
