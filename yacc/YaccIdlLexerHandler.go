package yacc

import (
	"github.com/bhbosman/yaccidl"
	"github.com/bhbosman/yaccpragma"
)

type idlLexerHandlerForTesting struct {
	pos  int
	data []lexem
}

func newIdlLexerHandlerForTesting(pos int, data []lexem) IYaccIdlLexerHandler {
	return &idlLexerHandlerForTesting{pos: pos, data: data}
}

func (y *idlLexerHandlerForTesting) ReadLexem() (*yaccidl.LexemValue, error) {
	for {
		y.pos++
		if y.pos == len(y.data)+1 {
			return nil, nil
		}
		switch y.data[y.pos-1].token {
		case yaccidl.Identifier:
			if v, ok := y.data[y.pos-1].data.(string); ok {
				return yaccidl.NewLexemStringValue(
					yaccidl.Identifier,
					nil,
					v,
					yaccidl.CurrentContext{},
					true)
			}
		case yaccidl.HashPragma:
			if v, ok := y.data[y.pos-1].data.(yaccpragma.IPragmaNode); ok {
				return yaccidl.NewLexemPragmaNodeValue(
					yaccidl.HashPragma,
					nil,
					v,
					yaccidl.CurrentContext{},
					true)
			}
		default:
			return yaccidl.NewLexemNoValue(
				y.data[y.pos-1].token,
				nil,
				yaccidl.CurrentContext{},
				true)
		}
	}
}
