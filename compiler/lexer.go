package compiler

import "strconv"

var charToTokenMap = map[rune]Token{
	'(':  T_OpeBracket,
	')':  T_CloBacket,
	'<':  T_ArrowLeft,
	'>':  T_ArrowRight,
	'\n': T_LineBreak,
}

func Lex(content string) []TokenMeta {

	lexemes := make([]TokenMeta, len(content))
	var line, col uint16 = 0, 0
	for i, run := range content {
		token, ok := charToTokenMap[run]
		if !ok {
			token = getDynamicToken(run)
		}
		col++

		lexemes[i] = *NewTokenMeta(token, SourcePosition{
			line,
			col,
			run,
		})

		if token == T_LineBreak {
			line++
			col = 0
		}
	}

	return lexemes
}

func getDynamicToken(r rune) Token {
	if isNumber(r) {
		return T_Number
	}

	return T_Letter
}

func isNumber(r rune) bool {
	if _, err := strconv.Atoi(string(r)); err == nil {
		return true
	}

	return false
}
