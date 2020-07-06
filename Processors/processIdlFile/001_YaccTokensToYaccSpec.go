package processIdlFile

import (
	"fmt"
	log2 "github.com/bhbosman/gocommon/log"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/yacc"
	"github.com/bhbosman/yaccidl"
	"log"
)

type YaccTokensToYaccSpec struct {
	logFactory *log2.LogFactory
	log        log2.SubSystemLogger
	ctx        *ctx2.GoYaccAppCtx
}

func (self YaccTokensToYaccSpec) Name() string {
	return self.log.Name()
}

func (self YaccTokensToYaccSpec) Run(input interface{}) (interface{}, error) {
	if lexer, ok := input.(yacc.IYaccIdlLexer); ok {
		self.log.LogWithLevel(0, func(logger *log.Logger) { logger.Print("Create parser...") })

		parser := yaccidl.YaccIdlNewParser()
		self.log.LogWithLevel(0, func(logger *log.Logger) {
			logger.Printf("Parse...")
		})
		v := parser.Parse(lexer)
		if v != 0 || lexer.ErrorOccurred() {
			return nil, fmt.Errorf("error reading idl file. Error: %v", lexer.ErrorMessage())
		}
		self.log.LogWithLevel(0, func(logger *log.Logger) {
			logger.Printf("Parse answer: %v", v)
		})
		specification := lexer.GetSpecification()
		if specification == nil {
			return nil, self.log.Error(fmt.Errorf("specification returned nil"))
		}
		return specification, nil
	}
	return nil, self.log.Error(fmt.Errorf("incorrect input type"))
}

func NewYaccTokensToYaccSpec(logFactory *log2.LogFactory, ctx *ctx2.GoYaccAppCtx) *YaccTokensToYaccSpec {
	return &YaccTokensToYaccSpec{
		logFactory: logFactory,
		log:        *logFactory.Create("YaccTokensToYaccSpec"),
		ctx:        ctx,
	}
}
