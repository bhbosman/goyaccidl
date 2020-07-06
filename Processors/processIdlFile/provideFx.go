package processIdlFile

import (
	"github.com/bhbosman/gocommon/Services/implementations"
	log2 "github.com/bhbosman/gocommon/log"
	"github.com/bhbosman/goyaccidl/Service"
	ctx2 "github.com/bhbosman/goyaccidl/ctx"
	"github.com/bhbosman/goyaccidl/process"
	"github.com/bhbosman/goyaccidl/typeHelpers"
	"github.com/bhbosman/goyaccidl/writers/gowriter"
	"github.com/bhbosman/goyaccidl/yaccSpecProcessing"
	yaccToken "github.com/bhbosman/yaccidl"
	"go.uber.org/fx"
)

func ProvideIdlProcess() fx.Option {
	return fx.Provide(
		func(params ProvideIdlProcessInParams) (ProvideIdlProcessOutParams, error) {
			//fileInformation := yaccToken.NewFileInformation()
			process, err := NewFileAllocationToGolang(
				params.LogFactory,
				params.Ctx,
				params.FileResolve,
				params.GoWriterFactory,
				params.IdlDefaultTypes,
				params.CreateKnownTypeService)
			if err != nil {
				return ProvideIdlProcessOutParams{}, err
			}
			return ProvideIdlProcessOutParams{
				YaccTokensToYaccSpec: NewYaccTokensToYaccSpec(
					params.LogFactory,
					params.Ctx),
				CreateLexer: NewCreateLexer(
					params.LogFactory,
					params.Ctx,
					params.IdlDefaultTypes,
					params.FileInformation),
				YaccSpecToDeclTypes: NewYaccSpecToDeclTypes(
					params.LogFactory,
					params.Ctx,
					params.UniqueSessionNumber,
					params.YaccSpecToDeclAll),
				FlattenDeclsToFileAllocation: NewFlattenDeclsToFileAllocation(
					params.LogFactory,
					params.Ctx,
					params.StructDclHelper,
					params.InterfaceDclHelper,
					params.TypeDefDclHelper,
					params.ConstantDclHelper,
					params.NativeDclHelper,
					params.ScopeDclHelper,
					params.PrimitiveDclHelper,
					params.UnionDclHelper,
					params.IdlDefaultTypes,
					params.IdlToGoTranslation,
					params.FileInformation,
					params.DclHelpers,
					params.CreateKnownTypeService,
					params.SequenceDcl,
					params.EnumDcl,
					params.RemoveTypedef,
					params.SeqStructDcl,
					params.ExpandInterfaces,
					params.StructStubDclHelper),
				FlattenDeclaredTypes: NewFlattenDeclaredTypes(
					params.LogFactory,
					params.Ctx,
					params.UniqueSessionNumber,
					params.FlattenStructDcl,
					params.FlattenInterfaceDcl,
					params.TypeDefDclHelper,
					params.FlattenConstantDcl,
					params.FlattenNativeDcl,
					params.ScopeDclHelper,
					params.PrimitiveDclHelper,
					params.SequenceDcl,
					params.UnionDcl,
					params.IdlDefaultTypes),
				FileAllocationToGolang: process,
				SaveImport:             NewSaveImport(params.LogFactory, params.Ctx),
			}, nil
		})
}

type ProvideIdlProcessInParams struct {
	fx.In
	LogFactory             *log2.LogFactory
	Ctx                    *ctx2.GoYaccAppCtx
	IdlDefaultTypes        *Service.IdlDefaultTypes
	UniqueSessionNumber    *implementations.UniqueSessionNumber
	StructDclHelper        *typeHelpers.StructDcl
	InterfaceDclHelper     *typeHelpers.InterfaceDcl
	TypeDefDclHelper       *typeHelpers.TypeDefDcl
	ConstantDclHelper      *typeHelpers.ConstantDcl
	NativeDclHelper        *typeHelpers.NativeDcl
	ScopeDclHelper         *typeHelpers.ScopeDcl
	SeqStructDcl           *typeHelpers.StructSequenceDcl
	PrimitiveDclHelper     *typeHelpers.PrimitiveDcl
	UnionDclHelper         *typeHelpers.UnionDcl
	FlattenStructDcl       *typeHelpers.StructDcl
	FlattenInterfaceDcl    *typeHelpers.InterfaceDcl
	FlattenConstantDcl     *typeHelpers.ConstantDcl
	FlattenNativeDcl       *typeHelpers.NativeDcl
	SequenceDcl            *typeHelpers.SequenceDcl
	UnionDcl               *typeHelpers.UnionDcl
	FileResolve            *ctx2.ResolveFileName
	GoWriterFactory        *gowriter.Factory
	YaccSpecToDeclAll      yaccSpecProcessing.IYaccSpecToDeclAll
	IdlToGoTranslation     *Service.IdlToGoTranslation
	DclHelpers             *Service.DclHelpers
	CreateKnownTypeService *CreateKnownTypeService
	EnumDcl                *typeHelpers.EnumDcl
	FileInformation        *yaccToken.FileInformation
	RemoveTypedef          *Service.RemoveTypedef
	ExpandInterfaces       *Service.ExpandInterfaces
	StructStubDclHelper    *typeHelpers.StructStubDcl
}

type ProvideIdlProcessOutParams struct {
	fx.Out
	CreateLexer                  process.IProcess `name:"CreateLexer"`
	YaccTokensToYaccSpec         process.IProcess `name:"YaccTokensToYaccSpec"`
	YaccSpecToDeclTypes          process.IProcess `name:"YaccSpecToDeclTypes"`
	FlattenDeclsToFileAllocation process.IProcess `name:"FlattenDeclsToFileAllocation"`
	FlattenDeclaredTypes         process.IProcess `name:"FlattenDeclaredTypes"`
	FileAllocationToGolang       process.IProcess `name:"FileAllocationToGolang"`
	SaveImport                   process.IProcess `name:"SaveImport"`
}
