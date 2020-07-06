package processIdlFile

import (
	"github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
	"go.uber.org/fx"
	log2 "log"
)

type Factory struct {
	ctx                         *ctx2.GoYaccAppCtx
	logger                      *log.SubSystemLogger
	loggerFactory               *log.LogFactory
	Lexer                       process.IProcess
	YaccTokensToYaccSpec        process.IProcess
	YaccSpecToDeclTypes         process.IProcess
	DeclTypesToFlattenDeclTypes process.IProcess
	FlattenDeclsToFileDecls     process.IProcess
	FileDeclToRenderGo          process.IProcess
	SaveImport                  process.IProcess
	idlDefaultTypes             *Service.IdlDefaultTypes
}

func (self Factory) Create() (process.IProcessWrapper, error) {
	self.logger.LogWithLevel(0, func(logger *log2.Logger) { logger.Println("Create FileAllocationToGolang Chain") })
	startProcess00, _ := process.NewProcessWrapper(self.loggerFactory, self.Lexer)
	processTokens01, _ := process.NewProcessWrapper(self.loggerFactory, self.YaccTokensToYaccSpec)
	processSpecification02, _ := process.NewProcessWrapper(self.loggerFactory, self.YaccSpecToDeclTypes)
	processFlatten03, _ := process.NewProcessWrapper(self.loggerFactory, self.DeclTypesToFlattenDeclTypes)
	flattenDeclsToFileDecls04, _ := process.NewProcessWrapper(self.loggerFactory, self.FlattenDeclsToFileDecls)
	renderToGo05, _ := process.NewProcessWrapper(self.loggerFactory, self.FileDeclToRenderGo)
	saveImport06, _ := process.NewProcessWrapper(self.loggerFactory, self.SaveImport)

	startProcess00.SetNext(processTokens01)
	processTokens01.SetNext(processSpecification02)
	processSpecification02.SetNext(processFlatten03)
	processFlatten03.SetNext(flattenDeclsToFileDecls04)
	flattenDeclsToFileDecls04.SetNext(renderToGo05)
	renderToGo05.SetNext(saveImport06)

	return startProcess00, nil
}

func NewProcessFactoryImpl(
	ctx *ctx2.GoYaccAppCtx,
	logger *log.SubSystemLogger,
	loggerFactory *log.LogFactory,
	Lexer process.IProcess,
	YaccTokensToYaccSpec process.IProcess,
	YaccSpecToDeclTypes process.IProcess,
	DeclTypesToFlattenDeclTypes process.IProcess,
	FlattenDeclsToFileDecls process.IProcess,
	FileDeclToRenderGo process.IProcess,
	SaveImport process.IProcess,
	idlDefaultTypes *Service.IdlDefaultTypes) (*Factory, error) {
	return &Factory{
		ctx:                         ctx,
		logger:                      logger,
		loggerFactory:               loggerFactory,
		Lexer:                       Lexer,
		YaccTokensToYaccSpec:        YaccTokensToYaccSpec,
		YaccSpecToDeclTypes:         YaccSpecToDeclTypes,
		DeclTypesToFlattenDeclTypes: DeclTypesToFlattenDeclTypes,
		FlattenDeclsToFileDecls:     FlattenDeclsToFileDecls,
		FileDeclToRenderGo:          FileDeclToRenderGo,
		SaveImport:                  SaveImport,
		idlDefaultTypes:             idlDefaultTypes,
	}, nil
}

func ProvideProcessFactory() fx.Option {
	return fx.Provide(
		func(params ProcessFactoryImplParams) (process.IFactory, error) {
			return NewProcessFactoryImpl(
				params.AppCtx,
				params.LoggerFactory.Create("ProcessFactory"),
				params.LoggerFactory,
				params.Lexer,
				params.YaccTokensToYaccSpec,
				params.YaccSpecToDeclTypes,
				params.DeclTypesToFlattenDeclTypes,
				params.FlattenDeclsToFileDecls,
				params.FileDeclToRenderGo,
				params.SaveImport,
				params.IdlDefaultTypes)
		})
}

type ProcessFactoryImplParams struct {
	fx.In
	AppCtx                      *ctx2.GoYaccAppCtx
	LoggerFactory               *log.LogFactory
	Lexer                       process.IProcess `name:"CreateLexer"`
	YaccTokensToYaccSpec        process.IProcess `name:"YaccTokensToYaccSpec"`
	YaccSpecToDeclTypes         process.IProcess `name:"YaccSpecToDeclTypes"`
	DeclTypesToFlattenDeclTypes process.IProcess `name:"FlattenDeclaredTypes"`
	FlattenDeclsToFileDecls     process.IProcess `name:"FlattenDeclsToFileAllocation"`
	FileDeclToRenderGo          process.IProcess `name:"FileAllocationToGolang"`
	SaveImport                  process.IProcess `name:"SaveImport"`
	IdlDefaultTypes             *Service.IdlDefaultTypes
}
