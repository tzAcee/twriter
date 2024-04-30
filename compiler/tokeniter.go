package compiler

type TokenIterator struct {
	tokens *[]TokenMeta
	index  uint16
}

func NewTokenIterator(t *[]TokenMeta) *TokenIterator {
	return &TokenIterator{t, 0}
}

func (t TokenIterator) getNoSkip() (TokenMeta, bool) {
	tokens := *t.tokens

	if t.index >= uint16(len(tokens)) {
		return TokenMeta{}, false
	}
	return tokens[t.index], true
}

func (t TokenIterator) Get() (TokenMeta, bool) {
	t.Skip()
	tokens := *t.tokens

	if t.index >= uint16(len(tokens)) {
		return TokenMeta{}, false
	}
	return tokens[t.index], true
}

func (t TokenIterator) Peek(n uint16) (TokenMeta, bool) {
	t.Skip()
	tokens := *t.tokens

	if t.index+n >= uint16(len(tokens)) {
		return TokenMeta{}, false
	}
	return tokens[t.index+n], true
}

func (t *TokenIterator) Skip() {
	for {
		if !t.isSkippable() {
			break
		}
		t.index += 1
	}
}

func (t TokenIterator) isSkippable() bool {
	tokenMeta, ok := t.getNoSkip()

	if !ok {
		return false
	}

	skippables := [1]Token{
		T_Whitespace,
	}

	for _, skipper := range skippables {
		if skipper == tokenMeta.token {
			return true
		}
	}

	return false

}

func (t *TokenIterator) Step(n uint16) {
	t.Skip()
	t.index += n
}
