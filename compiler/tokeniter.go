package compiler

type TokenIterator struct {
	tokens *[]TokenMeta
	index  uint16
}

func NewTokenIterator(t *[]TokenMeta) *TokenIterator {
	return &TokenIterator{t, 0}
}

func (t TokenIterator) Get() (TokenMeta, bool) {
	tokens := *t.tokens

	if t.index >= uint16(len(tokens)) {
		return TokenMeta{}, false
	}
	return tokens[t.index], true
}

func (t TokenIterator) Peek(n uint16) (TokenMeta, bool) {
	tokens := *t.tokens

	if t.index+n >= uint16(len(tokens)) {
		return TokenMeta{}, false
	}
	return tokens[t.index+n], true
}

func (t *TokenIterator) Step(n uint16) bool {
	if t.index+n >= uint16(len(*t.tokens)) {
		return false
	}

	t.index += n

	return true
}
