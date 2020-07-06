package processIdlFile

import (
	"github.com/bhbosman/gocommon"
	"github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/yacc"
	"github.com/bhbosman/lexidl"
	yaccToken "github.com/bhbosman/yaccidl"
)

type CreateLexer struct {
	logFactory      *log.LogFactory
	log             *log.SubSystemLogger
	ctx             *ctx2.GoYaccAppCtx
	idlDefaultTypes *Service.IdlDefaultTypes
	fileInformation *yaccToken.FileInformation
}

func (self CreateLexer) Run(in interface{}) (interface{}, error) {

	handler, err := lexidl.NewLexIdlHandler(
		self.ctx.InputFileName,
		self.ctx.IdlDefinitions.AssignFlags(),
		gocommon.NewByteReaderNoCloser(self.ctx.InputReader),
		self.fileInformation)

	if err != nil {
		return nil, err
	}
	return yacc.NewIdlLexer(
		self.idlDefaultTypes,
		self.log,
		handler)
}

func (self CreateLexer) Name() string {
	return self.log.Name()
}

func NewCreateLexer(
	logFactory *log.LogFactory,
	ctx *ctx2.GoYaccAppCtx,
	idlDefaultTypes *Service.IdlDefaultTypes,
	fileInformation *yaccToken.FileInformation) *CreateLexer {
	return &CreateLexer{
		logFactory:      logFactory,
		log:             logFactory.Create("CreateLexer"),
		ctx:             ctx,
		idlDefaultTypes: idlDefaultTypes,
		fileInformation: fileInformation,
	}
}
