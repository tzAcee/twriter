package compiler

type Token = uint8

const (
	T_OpeBracket Token = iota
	T_CloBacket
	T_ArrowLeft
	T_ArrowRight
	T_LineBreak
	T_Number
	T_Letter
)

type SourcePosition struct {
	line   uint16
	column uint16
	char   rune
}

type TokenMeta struct {
	token     Token
	sourcePos SourcePosition
}

func NewTokenMeta(t Token, sourcePos SourcePosition) *TokenMeta {
	return &TokenMeta{t, sourcePos}
}

func (t TokenMeta) Rune() rune {
	return t.sourcePos.char
}

func (t TokenMeta) Ln() uint16 {
	return t.sourcePos.line
}

func (t TokenMeta) Col() uint16 {
	return t.sourcePos.column
}
