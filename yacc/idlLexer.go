package yacc

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	"github.com/bhbosman/yaccidl"
	"log"
)

type IYaccIdlLexer interface {
	yaccidl.YaccIdlLexer
	GetSpecification() yaccidl.IYaccNode
	ErrorOccurred() bool
	ErrorMessage() string
}

type IYaccIdlLexerHandler interface {
	ReadLexem() (*yaccidl.LexemValue, error)
}

type IdlLexer struct {
	handler           IYaccIdlLexerHandler
	lexem             *yaccidl.LexemValue
	logger            *log2.SubSystemLogger
	specificationNode yaccidl.IYaccNode
	idlDefaultTypes   *Service.IdlDefaultTypes
	errorOccurred     bool
	errorMessage      string
}

func (self *IdlLexer) AssignSpecification(specificationNode yaccidl.IYaccNode) {
	self.specificationNode = specificationNode
}

func (self *IdlLexer) GetSpecification() yaccidl.IYaccNode {
	return self.specificationNode
}

func (self *IdlLexer) ErrorOccurred() bool {
	return self.errorOccurred
}

func (self *IdlLexer) ErrorMessage() string {
	return self.errorMessage
}

func (self *IdlLexer) ShortType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.ShortType()
}

func (self *IdlLexer) Int16Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Int16Type()
}

func (self *IdlLexer) Int32Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Int32Type()
}

func (self *IdlLexer) Int64Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Int64Type()
}

func (self *IdlLexer) UnsignedShortType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.UnsignedShortType()
}

func (self *IdlLexer) Uint16Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Uint16Type()
}

func (self *IdlLexer) Uint32Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Uint32Type()
}

func (self *IdlLexer) Uint64Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Uint64Type()
}

func (self *IdlLexer) LongType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.LongType()
}

func (self *IdlLexer) LongLongType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.LongLongType()
}

func (self *IdlLexer) UnsignedLongType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.UnsignedLongType()
}

func (self *IdlLexer) UnsignedLongLongType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.UnsignedLongLongType()
}

func (self *IdlLexer) CharType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.CharType()
}

func (self *IdlLexer) WideCharType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.WideCharType()
}

func (self *IdlLexer) StringType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.StringType()
}

func (self *IdlLexer) WideStringType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.WideStringType()
}

func (self *IdlLexer) BooleanType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.BooleanType()
}

func (self *IdlLexer) FloatType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.FloatType()
}

func (self *IdlLexer) DoubleType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.DoubleType()
}

func (self *IdlLexer) LongDoubleType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.LongDoubleType()
}

func (self *IdlLexer) VoidType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.VoidType()
}

func (self *IdlLexer) Int8Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.Int8Type()
}

func (self *IdlLexer) UInt8Type() yaccidl.IYaccNode {
	return self.idlDefaultTypes.UInt8Type()
}

func (self *IdlLexer) AnyType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.AnyType()
}

func (self *IdlLexer) ObjectType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.ObjectType()
}

func (self *IdlLexer) ValueBaseType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.ValueBaseType()
}

func (self *IdlLexer) OctetType() yaccidl.IYaccNode {
	return self.idlDefaultTypes.OctetType()
}

func (self *IdlLexer) LogWithLevel(level int, cb func(logger *log.Logger)) {
	self.logger.LogWithLevel(level, cb)
}

func (self *IdlLexer) Log(info string, params ...interface{}) {
	self.LogWithLevel(100, func(logger *log.Logger) {
		s := ""
		for _, p := range params {

			if s == "" {
				s = fmt.Sprintf("(%v)", p)
			} else {
				s = fmt.Sprintf("%v, (%v)", s, p)
			}
		}
		logger.Printf(" >>>%s [%v].<<<", info, s)
	})
}

func (self *IdlLexer) Lex(lval *yaccidl.YaccIdlSymType) int {
	for {
		localLexem, _ := self.handler.ReadLexem()
		if localLexem == nil {
			return 0
		}
		self.lexem = localLexem

		if self.lexem.Eof {
			return 0
		}

		if !self.lexem.ValidToken {
			continue
		}

		switch self.lexem.TypeKind {

		case yaccidl.WhiteSpace, yaccidl.SingleComment, yaccidl.None, yaccidl.MultiLineComment, yaccidl.HashLoadDefinition,
			yaccidl.HashDefine, yaccidl.HashUnDefine, yaccidl.Hashifdef, yaccidl.Hashifndef,
			yaccidl.Hashendif, yaccidl.Hashelse, yaccidl.HashInclude:
			continue
		case yaccidl.Integer_literal:
			lval.Value, _ = yaccidl.NewInt64Value(self.lexem.IntValue)
			return self.lexem.TypeKind
		case yaccidl.Identifier:
			lval.Node = yaccidl.NewYaccNode(self.lexem.StringValue, self.lexem)
			return self.lexem.TypeKind
		case yaccidl.RWtypeprefix:
			lval.Node = yaccidl.NewYaccNode(self.lexem.StringValue, self.lexem)
			return yaccidl.RWtypeprefix
		case yaccidl.HashPragma:
			continue
		default:
			lval.Node = yaccidl.NewYaccNode(self.lexem.StringValue, self.lexem)
			return self.lexem.TypeKind
		}
	}
}

func (d *IdlLexer) Error(s string) {
	d.errorOccurred = true
	d.errorMessage = s
	d.logger.LogWithLevel(0, func(logger *log.Logger) {
		logger.Printf("Error: %v. Lex: %v", s, d.lexem)
	})
}

func NewIdlLexer(
	idlDefaultTypes *Service.IdlDefaultTypes,
	logger *log2.SubSystemLogger,
	handler IYaccIdlLexerHandler) (IYaccIdlLexer, error) {
	if handler == nil {
		return nil, fmt.Errorf("handler not assigned")
	}

	return &IdlLexer{
		handler:         handler,
		logger:          logger,
		idlDefaultTypes: idlDefaultTypes,
	}, nil
}
